"use client";

import { useState } from "react";

const tabs = [
  { id: "overview", label: "Overview" },
  { id: "features", label: "Key Features" },
  { id: "getting-started", label: "Getting Started" },
  { id: "configuration", label: "Configuration Management" },
  { id: "faqs", label: "FAQs" },
  { id: "troubleshooting", label: "Troubleshooting" },
  { id: "support", label: "Contact Support" },
];

const DocumentationPage: React.FC = () => {
  const [activeTab, setActiveTab] = useState("overview");

  const renderContent = () => {
    switch (activeTab) {
      case "overview":
        return (
          <div>
            <h2 className="text-4xl font-semibold mb-4">Overview</h2>
            <p className="text-lg leading-relaxed">
              Welcome to the <strong>DevOps-as-a-Service Platform</strong>! This web-based platform
              simplifies infrastructure provisioning, CI/CD pipeline setup, and cloud cost estimation
              for startups and small teams.
            </p>
          </div>
        );
      case "features":
        return (
          <div>
            <h2 className="text-4xl font-semibold mb-4">Key Features</h2>
            <ul className="list-disc ml-6 text-lg leading-relaxed">
              <li>
                <strong>Infrastructure Automation:</strong> Automatically provision cloud resources
                using Terraform.
              </li>
              <li>
                <strong>CI/CD Pipeline Setup:</strong> Create CI/CD workflows using GitHub Actions.
              </li>
              <li>
                <strong>Cost Estimation Dashboard:</strong> Get real-time cost estimates for your
                infrastructure.
              </li>
              <li>
                <strong>Monitoring:</strong> Configure Prometheus and Grafana for metrics monitoring.
              </li>
            </ul>
          </div>
        );
      // Add other cases...
      default:
        return null;
    }
  };

  return (
    <div className="flex h-screen">
      {/* Sidebar */}
      <aside className="w-1/4 bg-white p-6 shadow-md border-r border-gray-200">
        <h2 className="text-2xl font-bold mb-6 text-gray-800">Documentation</h2>
        <ul className="space-y-4">
          {tabs.map((tab) => (
            <li key={tab.id}>
              <button
                className={`block w-full text-left text-lg px-4 py-3 rounded-lg transition-all duration-300 ${
                  activeTab === tab.id
                    ? "bg-blue-100 text-blue-600 font-semibold shadow-sm"
                    : "hover:bg-gray-100 text-gray-800"
                }`}
                onClick={() => setActiveTab(tab.id)}
              >
                {tab.label}
              </button>
            </li>
          ))}
        </ul>
      </aside>

      {/* Content Area */}
      <main className="w-3/4 p-8 bg-gray-50 overflow-y-auto">{renderContent()}</main>
    </div>
  );
};

export default DocumentationPage;
