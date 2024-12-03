"use client";

import { ConfigForm } from "@/components/generate/ConfigForm";

export default function GitHubActionsPage() {
  const fields = [
    {
      name: "workflowName",
      label: "Workflow Name",
      type: "text" as const,
      placeholder: "CI/CD Workflow",
      required: true,
    },
    {
      name: "triggerEvents",
      label: "Trigger Events",
      type: "text" as const,
      placeholder: "push, pull_request",
      required: true,
    },
    {
      name: "goVersion",
      label: "Go Version",
      type: "text" as const,
      placeholder: "1.18",
      required: true,
    },
    {
      name: "buildCommand",
      label: "Build Command",
      type: "text" as const,
      placeholder: "go build -v",
      required: true,
    },
    {
      name: "testCommand",
      label: "Test Command",
      type: "text" as const,
      placeholder: "go test -v",
      required: true,
    },
  ];

  const handleSubmit = () => {}
  // const handleSubmit = async (data: Record<string, string>) => {
  //   console.log("GitHub Actions config:", data);
    
  //   // You can make a request to your backend with these data, for example:
  //   const response = await fetch('/generate-github-action', {
  //     method: 'POST',
  //     headers: {
  //       'Content-Type': 'application/json',
  //     },
  //     body: JSON.stringify(data),
  //   });

  //   const yaml = await response.text();
  //   console.log("Generated YAML:", yaml);
  // };

  return (
    <div className="container px-4 py-6">
      <ConfigForm
        title="GitHub Actions Configuration"
        description="Configure GitHub Actions workflow settings"
        fields={fields}
        onSubmit={handleSubmit}
      />
    </div>
  );
}
