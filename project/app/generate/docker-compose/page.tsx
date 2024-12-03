"use client";

import { ConfigForm } from "@/components/generate/ConfigForm";

export default function DockerComposePage() {
  const fields = [
    {
      name: "serviceName",
      label: "Service Name",
      type: "text" as const,
      placeholder: "web",
      required: true,
    },
    {
      name: "image",
      label: "Image Name",
      type: "text" as const,
      placeholder: "my-app:latest",
      required: true,
    },
    {
      name: "ports",
      label: "Port Mapping",
      type: "text" as const,
      placeholder: "8080:80",
      required: true,
    },
    {
      name: "volumes",
      label: "Volumes",
      type: "textarea" as const,
      placeholder: "./data:/data",
      required: false,
    },
    {
      name: "environmentVariables",
      label: "Environment Variables",
      type: "textarea" as const,
      placeholder: "NODE_ENV=production\nAPI_URL=https://example.com",
      required: false,
    },
  ];

  const handleSubmit = (data: Record<string, string>) => {
    console.log("Docker Compose config:", data);
  };

  return (
    <div className="container px-4 py-6">
      <ConfigForm
        title="Docker Compose Configuration"
        description="Generate a Docker Compose file for your services"
        fields={fields}
        onSubmit={handleSubmit}
      />
    </div>
  );
}
