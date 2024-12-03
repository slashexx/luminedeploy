"use client";

import { ConfigForm } from "@/components/generate/ConfigForm";

export default function PrometheusPage() {
  const fields = [
    {
      name: "scrapeInterval",
      label: "Scrape Interval",
      type: "text" as const,
      placeholder: "15s",
      required: true,
    },
    {
      name: "evaluationInterval",
      label: "Evaluation Interval",
      type: "text" as const,
      placeholder: "15s",
      required: true,
    },
    {
      name: "targets",
      label: "Monitoring Targets",
      type: "textarea" as const,
      placeholder: "localhost:8080\nlocalhost:9090",
      required: true,
    },
    {
      name: "retention",
      label: "Data Retention",
      type: "select" as const,
      options: [
        { value: "15d", label: "15 days" },
        { value: "30d", label: "30 days" },
        { value: "60d", label: "60 days" },
      ],
      required: true,
    },
  ];

  const handleSubmit = (data: Record<string, string>) => {
    console.log("Prometheus config:", data);
  };

  return (
    <div className="container px-4 py-6">
      <ConfigForm
        title="Prometheus Configuration"
        description="Generate a Prometheus configuration file for monitoring"
        fields={fields}
        onSubmit={handleSubmit}
      />
    </div>
  );
}