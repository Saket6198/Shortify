"use client";

import { useState } from 'react';
import axios from 'axios';

export default function URLShortenerForm() {
  // State to store the URL input by the user, the shortened URL, and any error messages
  const [originalUrl, setOriginalUrl] = useState('');
  const [shortenedUrl, setShortenedUrl] = useState('');
  const [error, setError] = useState('');

  // Handle input change
  const handleInputChange = (e: any) => {
    setOriginalUrl(e.target.value);
  };

  // Handle form submission
  const handleSubmit = async (e:any) => {
    e.preventDefault();

    // Check if URL is empty
    if (!originalUrl) {
      setError('Please enter a URL');
      return;
    }

    try {
      // Make the POST request to the Go backend API
      const response = await axios.post('http://http://localhost:5173/shorten', {
        url: originalUrl,
      });

      // If a shortened URL is received in the response, set it to the state
      if (response.data.shortenedUrl) {
        setShortenedUrl(response.data.shortenedUrl);
        setError('');
      } else {
        setError('Failed to shorten URL');
      }
    } catch (err) {
      setError('Error while shortening URL');
      console.error(err);
    }
  };

  return (
    <div className="bg-black h-screen">
      <div className="container bg-white mx-auto flex flex-col items-center justify-center w-1/2 h-full p-8 rounded-lg shadow-lg">
        <p className="text-xl mb-4">
          Enter your URL to shorten it:
        </p>
        <form onSubmit={handleSubmit} className="w-full">
          <input
            type="text"
            placeholder="Enter URL"
            value={originalUrl}
            onChange={handleInputChange}
            className="border border-gray-300 p-2 rounded-md w-full mb-4"
          />
          <button type="submit" className="bg-blue-500 text-white p-2 rounded-md w-full">
            Shorten URL
          </button>
        </form>

        {/* Display any error messages */}
        {error && <div className="text-red-500 mt-4">{error}</div>}

        {/* Display the shortened URL if available */}
        {shortenedUrl && (
          <div className="mt-4">
            <p>Shortened URL:</p>
            <a href={shortenedUrl} target="_blank" rel="noopener noreferrer" className="text-blue-500">
              {shortenedUrl}
            </a>
          </div>
        )}
      </div>
    </div>
  );
}
