"use client";

export default function URLShortenerForm() {
  return (
    <div className="bg-black h-screen">
        <div className="container bg-white mx-auto flex flex-col items-center justify-center w-1/2 h-full p-8 rounded-lg shadow-lg">
            <p className="text-xl mb-4">
                Enter your URL to shorten it:
            </p>
            <form>
                <input type="text" placeholder="Enter URL" className="border border-gray-300 p-2 rounded-md w-full mb-4" />
                <button type="submit" className="bg-blue-500 text-white p-2 rounded-md w-full">Shorten URL</button>
            </form>
        </div>
    </div>
  );
}
