"use client"

import { useActionState } from "react"
import { useForm } from "react-hook-form"
import { zodResolver } from "@hookform/resolvers/zod"
import { urlShortenerSchema } from "../lib/validations/url-shortener-schema"
import urlAuth from "./actions/url-auth"
import { Form, FormControl, FormField, FormItem, FormMessage } from "@/components/ui/form"
import { Input } from "@/components/ui/input"
import { Button } from "@/components/ui/button"
import { Card, CardContent, CardDescription, CardFooter, CardHeader, CardTitle } from "@/components/ui/card"
import { Alert, AlertDescription, AlertTitle } from "@/components/ui/alert"
import { Loader2, LinkIcon, Copy, ExternalLink, AlertCircle, CheckCircle2 } from "lucide-react"
import { useState } from "react"

export default function Landing() {
  const [result, formAction, isPending] = useActionState(urlAuth, undefined)

  const [copied, setCopied] = useState(false)

  const form = useForm({
    resolver: zodResolver(urlShortenerSchema),
    defaultValues: {
      url: "",
    },
  })

  const copyToClipboard = (text: string) => {
    navigator.clipboard.writeText(text)
    setCopied(true)
    setTimeout(() => setCopied(false), 2000)
  }

  // Extract error message from Zod validation errors
  const getErrorMessage = () => {
    if (!result) return null

    if ("message" in result) {
      return result.message
    }

    if ("url" in result && result.url && typeof result.url === "object" && "_errors" in result.url) {
      return (result.url._errors as string[])[0]
    }

    if ("error" in result) {
      return result.error
    }

    return null
  }

  const errorMessage = getErrorMessage()

  return (
    <div className="min-h-screen bg-gradient-to-b from-slate-50 to-slate-100 flex flex-col items-center justify-center p-4">
      <div className="w-full max-w-md">
        <Card className="border-slate-200 shadow-lg">
          <CardHeader className="text-center">
            <div className="mx-auto bg-slate-100 p-3 rounded-full w-16 h-16 flex items-center justify-center mb-2">
              <LinkIcon className="h-8 w-8 text-slate-700" />
            </div>
            <CardTitle className="text-2xl font-bold text-slate-800">URL Shortener</CardTitle>
            <CardDescription className="text-slate-500">
              Transform your long URLs into short, shareable links
            </CardDescription>
          </CardHeader>
          <CardContent>
            <Form {...form}>
              <form action={formAction} className="space-y-4">
                <FormField
                  control={form.control}
                  name="url"
                  render={({ field }) => (
                    <FormItem>
                      <FormControl>
                        <div className="relative">
                          <Input
                            {...field}
                            placeholder="https://your-long-url.com"
                            disabled={isPending}
                            className="pr-10 h-12 border-slate-300 focus:border-slate-500 focus:ring-slate-500"
                          />
                          {field.value && (
                            <button
                              type="button"
                              className="absolute right-3 top-1/2 -translate-y-1/2 text-slate-400 hover:text-slate-600"
                              onClick={() => field.onChange("")}
                            >
                              ×
                            </button>
                          )}
                        </div>
                      </FormControl>
                      <FormMessage />
                    </FormItem>
                  )}
                />
                <Button
                  type="submit"
                  disabled={isPending}
                  className="w-full h-12 bg-slate-800 hover:bg-slate-700 text-white"
                >
                  {isPending ? (
                    <>
                      <Loader2 className="mr-2 h-4 w-4 animate-spin" />
                      Shortening...
                    </>
                  ) : (
                    "Shorten URL"
                  )}
                </Button>
              </form>
            </Form>

            {errorMessage && (
              <Alert variant="destructive" className="mt-4">
                <AlertCircle className="h-4 w-4" />
                <AlertTitle>Error</AlertTitle>
                <AlertDescription>{errorMessage}</AlertDescription>
              </Alert>
            )}

            {result && "shortenedUrl" in result && (
              <div className="mt-4 space-y-3">
                <Alert variant="default" className="bg-green-50 border-green-200">
                  <CheckCircle2 className="h-4 w-4 text-green-600" />
                  <AlertTitle className="text-green-800">Success!</AlertTitle>
                  <AlertDescription className="text-green-700">
                    Your URL has been shortened successfully.
                  </AlertDescription>
                </Alert>

                <div className="relative mt-2 p-3 bg-slate-50 rounded-md border border-slate-200">
                  <p className="text-sm font-medium text-slate-600 mb-1">Your shortened URL:</p>
                  <div className="flex items-center">
                    <a
                      href={result.shortenedUrl}
                      target="_blank"
                      rel="noopener noreferrer"
                      className="text-blue-600 hover:text-blue-800 text-sm font-medium truncate mr-2"
                    >
                      {result.shortenedUrl}
                    </a>
                    <div className="flex-shrink-0 ml-auto space-x-1">
                      <Button
                        size="sm"
                        variant="outline"
                        className="h-8 px-2"
                        onClick={() => result.shortenedUrl && copyToClipboard(result.shortenedUrl)}
                      >
                        {copied ? <CheckCircle2 className="h-4 w-4 text-green-600" /> : <Copy className="h-4 w-4" />}
                      </Button>
                      <Button
                        size="sm"
                        variant="outline"
                        className="h-8 px-2"
                        onClick={() => window.open(result.shortenedUrl, "_blank")}
                      >
                        <ExternalLink className="h-4 w-4" />
                      </Button>
                    </div>
                  </div>
                </div>
              </div>
            )}
          </CardContent>
          <CardFooter className="flex justify-center border-t border-slate-100 pt-4">
            <p className="text-xs text-slate-500">Shorten URLs quickly and securely</p>
          </CardFooter>
        </Card>
      </div>
    </div>
  )
}
