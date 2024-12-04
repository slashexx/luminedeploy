'use client';

import { useEffect, useState } from 'react';
import { FileData } from '@/app/types/files';
import { FileList } from '@/components/files/FileList';
import { LoadingSpinner } from '@/components/ui/loading-spinner';

export default function UploadFilesPage() {
  const [files, setFiles] = useState<FileData[]>([]);
  const [loading, setLoading] = useState<boolean>(true);
  
  useEffect(() => {
    const fetchFiles = async () => {
      try {
        const response = await fetch('http://localhost:8080/files');
        if (!response.ok) throw new Error('Failed to fetch files');
        const data = await response.json();
        setFiles(data);
      } catch (error) {
        console.error('Error fetching files:', error);
      } finally {
        setLoading(false);
      }
    };

    fetchFiles();
  }, []);

  const handleDeployOnAWS = () => {
    console.log('Deploying on AWS...');
  };

  const handleGenerateZip = () => {
    console.log('Generating zip file...');
  };

  return (
    <div className="min-h-screen p-8">
      <div className="max-w-4xl mx-auto flex flex-col min-h-[600px]">
        <h2 className="text-2xl font-semibold mb-6">Uploaded Files</h2>
        
        <div className="flex-grow">
          {loading ? (
            <LoadingSpinner />
          ) : (
            <FileList files={files} />
          )}
        </div>

        {!loading && (
          <div className="grid grid-cols-2 gap-6 mt-8">
            <button
              onClick={handleDeployOnAWS}
              className="py-4 px-6 bg-black text-white rounded-lg hover:bg-gray-800 transition-colors duration-200 text-lg font-medium"
            >
              Deploy on AWS
            </button>
            <button
              onClick={handleGenerateZip}
              className="py-4 px-6 bg-black text-white rounded-lg hover:bg-gray-800 transition-colors duration-200 text-lg font-medium"
            >
              Generate Zip File
            </button>
          </div>
        )}
      </div>
    </div>
  );
}