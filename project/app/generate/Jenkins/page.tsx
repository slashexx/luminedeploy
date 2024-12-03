"use client";

import { ConfigForm } from "@/components/generate/ConfigForm";

export default function JenkinsPage() {
  const fields = [
    {
      name: "jobName",
      label: "Job Name",
      type: "text" as const,
      placeholder: "my-jenkins-job",
      required: true,
    },
    {
      name: "buildTrigger",
      label: "Build Trigger",
      type: "text" as const,
      placeholder: "GitHub Webhook",
      required: true,
    },
    {
      name: "buildCommand",
      label: "Build Command",
      type: "text" as const,
      placeholder: "npm run build",
      required: true,
    },
  ];

  const handleSubmit = (data: Record<string, string>) => {
    console.log("Jenkins config:", data);
  };

  return (
    <div className="container px-4 py-6">
      <ConfigForm
        title="Jenkins Configuration"
        description="Configure Jenkins job settings"
        fields={fields}
        onSubmit={handleSubmit}
      />
    </div>
  );
}
