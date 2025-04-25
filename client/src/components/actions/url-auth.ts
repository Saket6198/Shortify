"use server";

import { urlShortenerSchema } from "@/lib/validations/url-shortener-schema";
import axios from "axios";
import { ZodError } from "zod";

export default async function urlAuth(_: any, formData: FormData){
    try{
        const rawUrl = await urlShortenerSchema.parseAsync({
            url: formData.get("url")
        });

        const res = await axios.post("http://localhost:5000/shorten",{
            url: rawUrl.url
        });

        return {
            success: true,
            shortenedUrl: `http://localhost:5000/redirect/${res.data.shortened_url}`
        };
    }catch(err){
        if (err instanceof ZodError){
            return err.format();
        }
        if(axios.isAxiosError(err)){
            return {
                message: err.response?.data.message || "An error occurred while shortening the URL."
            }
        }
        return { error: "Unknown error" };
    }
}