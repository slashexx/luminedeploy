"use client";

import React from "react";
import { Upload, Loader2 } from "lucide-react";
import { Container } from "@/components/ui/container";
import { Progress } from "@/components/ui/progress";
import { useRouter } from "next/router";
// import axios from "axios";

export default class UploadZone extends React.Component {
  state = {
    isUploading: false,
    progress: 0,
  };

  router = useRouter();

  onDrop = async (acceptedFiles: File[]) => {
    this.setState({ isUploading: true, progress: 0 });

    const formData = new FormData();
    acceptedFiles.forEach((file) => formData.append("file", file)); // Correct key is 'file'

    console.log("Accepted Files: ", acceptedFiles);

    try {
      const response = await fetch("http://localhost:8080/upload", { // Replace with your backend URL
        method: "POST",
        body: formData,
      });
      if (response.ok) {
        console.log("Request has been sent");

        window.location.href = "/files";
      } else {
        throw new Error("Upload failed. Please try again.");
      }
    } catch (error) {
      console.error("Error uploading file:", error);
    } finally {
      this.setState({ isUploading: false, progress: 100 }); // Simulate completion
    }
  };

  render() {
    const { isUploading, progress } = this.state;

    return (
      <Container className="py-12">
        <div
          className={`relative flex min-h-[300px] cursor-pointer flex-col items-center justify-center rounded-lg border-2 border-dashed bg-muted/50 p-12 text-center transition-colors ${isUploading ? "border-primary" : "border-muted-foreground/25"}`}
        >
          <input
            type="file"
            onChange={(e) => this.onDrop(Array.from(e.target.files || []))}  // Convert FileList to File[]
            disabled={isUploading}
            multiple
          />
          {isUploading ? (
            <div className="flex flex-col items-center gap-4">
              <Loader2 className="h-12 w-12 animate-spin text-primary" />
              <div className="w-64">
                <Progress value={progress} className="h-2" />
              </div>
              <p className="text-sm text-muted-foreground">
                Uploading your project...
              </p>
            </div>
          ) : (
            <>
              <Upload className="mb-4 h-12 w-12 text-muted-foreground" />
              <h3 className="mb-2 text-lg font-medium">
                Drop your project folder here
              </h3>
              <p className="text-sm text-muted-foreground">
                or click to select the folder you want to deploy
              </p>
            </>
          )}
        </div>
      </Container>
    );
  }
}
