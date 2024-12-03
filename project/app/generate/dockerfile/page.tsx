"use client";

import { ConfigForm } from "@/components/generate/ConfigForm";

export default function DockerfilePage() {
  const fields = [
    {
      name: "baseImage",
      label: "Base Image",
      type: "text" as const,
      placeholder: "node:14",
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
      placeholder: "RUN npm install",
      required: true,
    },
    {
      name: "startCommand",
      label: "Start Command",
      type: "text" as const,
      placeholder: "CMD [\"npm\", \"start\"]",
      required: true,
    },
  ];

  const handleSubmit = (data: Record<string, string>) => {
    console.log("Dockerfile config:", data);
  };

  return (
    <div className="container px-4 py-6">
      <ConfigForm
        title="Dockerfile Configuration"
        description="Generate a Dockerfile for your application"
        fields={fields}
        onSubmit={handleSubmit}
      />
    </div>
  );
}
