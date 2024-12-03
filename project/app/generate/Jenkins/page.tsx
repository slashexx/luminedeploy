"use client";

import { ConfigForm } from "@/components/generate/ConfigForm";
import { useState } from "react";

export default function JenkinsPage() {
  const [formData, setFormData] = useState({
    pipelineName: "DefaultPipeline",
    branchName: "main",
    buildCommand: "go build",
    testCommand: "go test ./...",
    agentLabel: "linux",
  });

  const fields = [
    {
      name: "pipelineName",
      label: "Pipeline Name",
      type: "text" as const,
      placeholder: "DefaultPipeline",
      required: true,
    },
    {
      name: "branchName",
      label: "Branch Name",
      type: "text" as const,
      placeholder: "main",
      required: true,
    },
    {
      name: "buildCommand",
      label: "Build Command",
      type: "text" as const,
      placeholder: "go build",
      required: true,
    },
    {
      name: "testCommand",
      label: "Test Command",
      type: "text" as const,
      placeholder: "go test ./...",
      required: true,
    },
    {
      name: "agentLabel",
      label: "Agent Label",
      type: "text" as const,
      placeholder: "linux",
      required: true,
    },
  ];

  const handleSubmit = () => {}
  // const handleSubmit = async (data: Record<string, string>) => {
  //   setFormData(data);

  //   // Send the form data to the backend
  //   const response = await fetch("/generate-jenkinsfile", {
  //     method: "GET",
  //     headers: {
  //       "Content-Type": "application/json",
  //     },
  //     body: JSON.stringify(data),
  //   });

  //   const jenkinsfile = await response.text();
  //   console.log("Generated Jenkinsfile:", jenkinsfile);
  // };

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
