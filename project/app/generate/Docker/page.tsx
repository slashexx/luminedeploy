"use client";

import { ConfigForm } from "@/components/generate/ConfigForm";

export default function DockerPage() {
  const fields = [
    {
      name: "imageName",
      label: "Docker Image Name",
      type: "text" as const,
      placeholder: "my-app",
      required: true,
    },
    {
      name: "tag",
      label: "Tag",
      type: "text" as const,
      placeholder: "latest",
      required: true,
    },
    {
      name: "port",
      label: "Port Mapping",
      type: "text" as const,
      placeholder: "8080:80",
      required: true,
    },
  ];

  const handleSubmit = (data: Record<string, string>) => {
    console.log("Docker config:", data);
  };

  return (
    <div className="container px-4 py-6">
      <ConfigForm
        title="Docker Configuration"
        description="Configure Docker container settings"
        fields={fields}
        onSubmit={handleSubmit}
      />
    </div>
  );
}
