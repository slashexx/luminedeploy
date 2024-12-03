"use client";

import { Card } from "@/components/ui/card";
import { Button } from "@/components/ui/button";
import Link from "next/link";
import {
  Github,
  
  Cloud,
  BarChart,
  
  Database,
  Monitor,
  Settings,
} from "lucide-react";
import { ConfigCard } from "@/components/generate/ConfigCard";

export default function GenerateConfigsPage() {
  return (
    <div className="container px-4 py-6">
      <h1 className="text-3xl font-bold tracking-tight mb-2">
        Generate Configs at One Click
      </h1>
      <p className="text-muted-foreground mb-8">
        Choose a service to generate configuration files
      </p>

      <div className="grid gap-6 md:grid-cols-2 lg:grid-cols-3">
        <ConfigCard
          title="CI/CD"
          description="Configure your continuous integration and deployment pipelines"
          icon={Settings}
          items={[
            { name: "GitHub Actions", path: "/generate/github-actions", icon: Github },
            // { name: "Jenkins", path: "/generate/jenkins", icon: Jenkins },
          ]}
        />

        <ConfigCard
          title="Cloud Provider"
          description="Set up your cloud infrastructure and services"
          icon={Cloud}
          items={[
            { name: "AWS", path: "/generate/aws" },
            { name: "Azure", path: "/generate/azure" },
            { name: "GCP", path: "/generate/gcp" },
          ]}
        />

        <ConfigCard
          title="Monitoring"
          description="Configure monitoring and observability tools"
          icon={Monitor}
          items={[
            { name: "Prometheus", path: "/generate/prometheus" },
            { name: "Grafana", path: "/generate/grafana" },
          ]}
        />

        <ConfigCard
          title="Docker"
          description="Generate Docker and container orchestration configs"
          icon={Github}
          items={[
            { name: "Dockerfile", path: "/generate/dockerfile" },
            { name: "Docker Compose", path: "/generate/docker-compose" },
          ]}
        />
      </div>
    </div>
  );
}