"use client";

import { useState } from "react";
import { ConfigForm } from "@/components/generate/ConfigForm";
import { Value } from "@radix-ui/react-select";

export default function DockerfilePage() {
  const fields = [
    {
      name: "baseImage",
      label: "Base Image",
      type: "text" as const,
      Value: "golang:1.22",
    },
    {
      name: "workingDirectory",
      label: "Working Directory",
      type: "text" as const,
      Value: "/app",
    },
    {
      name: "copyCommand",
      label: "Copy Files Command",
      type: "text" as const,
      Value: "COPY . .",
    },
    {
      name: "installCommand",
      label: "Install Dependencies Command",
      type: "text" as const,
      Value: "RUN go mod tidy && go build -o app",
    },
    {
      name: "startCommand",
      label: "Start Command",
      type: "text" as const,
      Value: 'CMD ["./app"]',
    },
  ];

  const [responseMessage, setResponseMessage] = useState("");
  const [errorMessage, setErrorMessage] = useState("");

  const handleSubmit = async (data: Record<string, string>) => {
    setResponseMessage(""); // Clear any previous messages
    setErrorMessage("");

    try {
      const res = await fetch(
        "http://localhost:8080/api/generate-go-dockerfile",
        {
          method: "POST",
          headers: {
            "Content-Type": "application/json",
          },
          body: JSON.stringify(data),
        }
      );

      // Check if the response is JSON
      if (res.ok) {
        const result = await res.json();
        setResponseMessage(
          result.message || "Dockerfile generated successfully!"
        );
      } else {
        // Try to parse error message if returned as JSON
        try {
          const errorResult = await res.json();
          setErrorMessage(
            `Error: ${errorResult.message || "Failed to generate Dockerfile"}`
          );
        } catch {
          // If not JSON, use plain text or status
          setErrorMessage(`Error: ${res.status} - ${res.statusText}`);
        }
      }
    } catch (error) {
      console.error("Error generating Dockerfile:", error);
      setErrorMessage(
        "An unexpected error occurred while generating the Dockerfile."
      );
    }
  };

  return (
    <div className="container px-4 py-6">
      <ConfigForm
        title="Dockerfile Configuration"
        description="Generate a Dockerfile for your application"
        fields={fields}
        onSubmit={handleSubmit}
      />
      {responseMessage && (
        <div className="mt-4 text-green-600">{responseMessage}</div>
      )}
      {errorMessage && <div className="mt-4 text-red-600">{errorMessage}</div>}
    </div>
  );
}
