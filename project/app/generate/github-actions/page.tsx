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
      name: "branch",
      label: "Branch",
      type: "text" as const,
      placeholder: "main",
      required: true,
    },
    {
      name: "nodeVersion",
      label: "Node.js Version",
      type: "text" as const,
      placeholder: "14.x",
      required: true,
    },
  ];

  const handleSubmit = (data: Record<string, string>) => {
    console.log("GitHub Actions config:", data);
  };

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
