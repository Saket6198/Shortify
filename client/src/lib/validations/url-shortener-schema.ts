import {object, string} from "zod";

export const urlShortenerSchema = object({
    url: string({required_error: "URL is required"})
        .url({message: "Invalid URL"})
        .min(1, {message: "URL is required"})
        .max(2048, {message: "URL must be less than 2048 characters"})
});
