"use client";

import { Button } from "@/components/ui/button";
import { Card } from "@/components/ui/card";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { Textarea } from "@/components/ui/textarea";
import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from "@/components/ui/select";

interface ConfigField {
  name: string;
  label: string;
  type: "text" | "textarea" | "select";
  placeholder?: string;
  options?: { value: string; label: string }[];
  required?: boolean;
}

interface ConfigFormProps {
  title: string;
  description: string;
  fields: ConfigField[];
  onSubmit: (data: Record<string, string>) => void;
}

export function ConfigForm({ title, description, fields, onSubmit }: ConfigFormProps) {
  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    const formData = new FormData(e.target as HTMLFormElement);
    const data: Record<string, string> = {};
    fields.forEach((field) => {
      data[field.name] = formData.get(field.name) as string;
    });
    onSubmit(data);
  };

  return (
    <Card className="p-6">
      <h2 className="text-2xl font-bold mb-2">{title}</h2>
      <p className="text-muted-foreground mb-6">{description}</p>
      
      <form onSubmit={handleSubmit} className="space-y-6">
        {fields.map((field) => (
          <div key={field.name} className="space-y-2">
            <Label htmlFor={field.name}>{field.label}</Label>
            {field.type === "textarea" ? (
              <Textarea
                id={field.name}
                name={field.name}
                placeholder={field.placeholder}
                required={field.required}
              />
            ) : field.type === "select" ? (
              <Select name={field.name} required={field.required}>
                <SelectTrigger>
                  <SelectValue placeholder={field.placeholder} />
                </SelectTrigger>
                <SelectContent>
                  {field.options?.map((option) => (
                    <SelectItem key={option.value} value={option.value}>
                      {option.label}
                    </SelectItem>
                  ))}
                </SelectContent>
              </Select>
            ) : (
              <Input
                id={field.name}
                name={field.name}
                type="text"
                placeholder={field.placeholder}
                required={field.required}
              />
            )}
          </div>
        ))}
        <Button type="submit" className="w-full">Generate Configuration</Button>
      </form>
    </Card>
  );
}