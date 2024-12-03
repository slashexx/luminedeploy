"use client";

import { useState } from "react";
import { ConfigForm } from "@/components/generate/ConfigForm";

export default function DockerfilePage() {
  const fields = [
    {
      name: "baseImage",
      label: "Base Image",
      type: "text" as const,
      placeholder: "golang:1.22",
      required: true,
    },
    {
      name: "workingDirectory",
      label: "Working Directory",
      type: "text" as const,
      placeholder: "/app",
      required: true,
    },
    {
      name: "copyCommand",
      label: "Copy Files Command",
      type: "text" as const,
      placeholder: "COPY . .",
      required: true,
    },
    {
      name: "installCommand",
      label: "Install Dependencies Command",
      type: "text" as const,
      placeholder: "RUN go mod tidy && go build -o app",
      required: true,
    },
    {
      name: "startCommand",
      label: "Start Command",
      type: "text" as const,
      placeholder: "CMD [\"./app\"]",
      required: true,
    },
  ];

  const [responseMessage, setResponseMessage] = useState("");


  const handleSubmit = () => {}
  // const handleSubmit = async (data: Record<string, string>) => {
  //   try {
  //     const res = await fetch("/api/generate-dockerfile", {
  //       method: "POST",
  //       headers: {
  //         "Content-Type": "application/json",
  //       },
  //       body: JSON.stringify(data),
  //     });

  //     const result = await res.json();
  //     if (res.ok) {
  //       setResponseMessage(result.message);
  //     } else {
  //       setResponseMessage(`Error: ${result.message}`);
  //     }
  //   } catch (error) {
  //     console.error("Error generating Dockerfile:", error);
  //     setResponseMessage("An error occurred while generating the Dockerfile.");
  //   }
  // };

  return (
    <div className="container px-4 py-6">
      <ConfigForm
        title="Dockerfile Configuration"
        description="Generate a Dockerfile for your application"
        fields={fields}
        onSubmit={handleSubmit}
      />
      {responseMessage && <div className="mt-4">{responseMessage}</div>}
    </div>
  );
}
