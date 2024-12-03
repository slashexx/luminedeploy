"use client";

import { Button } from "@/components/ui/button";
import { Card } from "@/components/ui/card";
import { Tabs, TabsContent, TabsList, TabsTrigger } from "@/components/ui/tabs";
import { BarChart3, Cloud, Shield, Zap } from "lucide-react";
import DashboardMetrics from "@/components/dashboard/DashboardMetrics";
import CostOverview from "@/components/dashboard/CostOverview";
import DeploymentStatus from "@/components/dashboard/DeploymentStatus";
import SecurityOverview from "@/components/dashboard/SecurityOverview";

export default function Home() {
  return (
    <main className="min-h-screen bg-background">
      <nav className="border-b">
        <div className="container flex h-16 items-center px-4">
          <div className="flex items-center space-x-2">
            <Zap className="h-6 w-6 text-primary" />
            <span className="text-xl font-bold">Lumine</span>
          </div>
          <div className="ml-auto flex items-center space-x-4">
            <Button variant="ghost">Documentation</Button>
            <Button variant="ghost">Support</Button>
            <Button>Get Started</Button>
          </div>
        </div>
      </nav>

      <div className="container px-4 py-6">
        <div className="mb-8">
          <h1 className="text-3xl font-bold tracking-tight">Welcome to Lumine</h1>
          <p className="text-muted-foreground">
            Your all-in-one DevOps automation platform
          </p>
        </div>

        <div className="grid gap-6 md:grid-cols-4">
          <Card className="p-4">
            <div className="flex items-center space-x-2">
              <Cloud className="h-5 w-5 text-primary" />
              <span className="font-medium">Infrastructure</span>
            </div>
            <div className="mt-3">
              <div className="text-2xl font-bold">12</div>
              <div className="text-sm text-muted-foreground">Active Services</div>
            </div>
          </Card>
          <Card className="p-4">
            <div className="flex items-center space-x-2">
              <BarChart3 className="h-5 w-5 text-primary" />
              <span className="font-medium">Cost</span>
            </div>
            <div className="mt-3">
              <div className="text-2xl font-bold">$2,451</div>
              <div className="text-sm text-muted-foreground">Monthly Spend</div>
            </div>
          </Card>
          <Card className="p-4">
            <div className="flex items-center space-x-2">
              <Shield className="h-5 w-5 text-primary" />
              <span className="font-medium">Security</span>
            </div>
            <div className="mt-3">
              <div className="text-2xl font-bold">98%</div>
              <div className="text-sm text-muted-foreground">Security Score</div>
            </div>
          </Card>
          <Card className="p-4">
            <div className="flex items-center space-x-2">
              <Zap className="h-5 w-5 text-primary" />
              <span className="font-medium">Performance</span>
            </div>
            <div className="mt-3">
              <div className="text-2xl font-bold">99.9%</div>
              <div className="text-sm text-muted-foreground">Uptime</div>
            </div>
          </Card>
        </div>

        <Tabs defaultValue="metrics" className="mt-6">
          <TabsList>
            <TabsTrigger value="metrics">Metrics</TabsTrigger>
            <TabsTrigger value="costs">Costs</TabsTrigger>
            <TabsTrigger value="deployments">Deployments</TabsTrigger>
            <TabsTrigger value="security">Security</TabsTrigger>
          </TabsList>
          <TabsContent value="metrics">
            <DashboardMetrics />
          </TabsContent>
          <TabsContent value="costs">
            <CostOverview />
          </TabsContent>
          <TabsContent value="deployments">
            <DeploymentStatus />
          </TabsContent>
          <TabsContent value="security">
            <SecurityOverview />
          </TabsContent>
        </Tabs>
      </div>
    </main>
  );
}