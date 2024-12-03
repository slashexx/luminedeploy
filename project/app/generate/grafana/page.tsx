import React from "react";

interface Field {
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
    fields: Field[];
    onSubmit: (data: Record<string, string>) => void;
}

export const ConfigForm: React.FC<ConfigFormProps> = ({
    title,
    description,
    fields,
    onSubmit,
}) => {
    const [formData, setFormData] = React.useState<Record<string, string>>({});

    const handleChange = (e: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement | HTMLSelectElement>) => {
        const { name, value } = e.target;
        setFormData({ ...formData, [name]: value });
    };

    const handleSubmit = (e: React.FormEvent) => {
        e.preventDefault();
        onSubmit(formData);
    };

    return (
        <form onSubmit={handleSubmit} className="space-y-6">
            <h1 className="text-2xl font-semibold text-gray-800">{title}</h1>
            <p className="text-gray-600">{description}</p>

            {fields.map((field) => (
                <div key={field.name} className="flex flex-col">
                    <label htmlFor={field.name} className="text-sm font-medium text-gray-700">
                        {field.label}
                    </label>
                    {field.type === "textarea" ? (
                        <textarea
                            id={field.name}
                            name={field.name}
                            placeholder={field.placeholder}
                            required={field.required}
                            className="mt-1 p-3 border border-gray-300 rounded-lg shadow-sm focus:ring-blue-500 focus:border-blue-500"
                            onChange={handleChange}
                        />
                    ) : field.type === "select" ? (
                        <select
                            id={field.name}
                            name={field.name}
                            required={field.required}
                            className="mt-1 p-3 border border-gray-300 rounded-lg shadow-sm focus:ring-blue-500 focus:border-blue-500"
                            onChange={handleChange}
                        >
                            {field.options?.map((option) => (
                                <option key={option.value} value={option.value}>
                                    {option.label}
                                </option>
                            ))}
                        </select>
                    ) : (
                        <input
                            id={field.name}
                            name={field.name}
                            type="text"
                            placeholder={field.placeholder}
                            required={field.required}
                            className="mt-1 p-3 border border-gray-300 rounded-lg shadow-sm focus:ring-blue-500 focus:border-blue-500"
                            onChange={handleChange}
                        />
                    )}
                </div>
            ))}

            <button
                type="submit"
                className="w-full py-3 bg-blue-600 text-white font-medium rounded-lg hover:bg-blue-700 focus:ring-4 focus:ring-blue-300"
            >
                Generate Config
            </button>
        </form>
    );
};