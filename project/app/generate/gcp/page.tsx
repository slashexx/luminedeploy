"use client";

import { ConfigForm } from "@/components/generate/ConfigForm";

export default function GCPPage() {
    const fields = [
        {
            name: "projectId",
            label: "Project ID",
            type: "text" as const,
            placeholder: "my-gcp-project",
            required: true,
        },
        {
            name: "serviceAccountKey",
            label: "Service Account Key",
            type: "textarea" as const,
            placeholder: "Paste your service account key JSON here",
            required: true,
        },
        {
            name: "region",
            label: "Region",
            type: "text" as const,
            placeholder: "us-central1",
            required: true,
        },
    ];

    const handleSubmit = (data: Record<string, string>) => {
        console.log("GCP config:", data);
    };

    return (
        <div className="container px-4 py-6">
            <ConfigForm
                title="GCP Configuration"
                description="Configure GCP settings for deployment"
                fields={fields}
                onSubmit={handleSubmit}
            />
        </div>
    );
}
