import { useEffect, useState } from "react";
import { useTranslation } from "react-i18next";
import { Form, Input, type FormInstance } from "antd";
import { createSchemaFieldRule } from "antd-zod";
import { z } from "zod";

import { type BytePlusAccessConfig } from "@/domain/access";

type AccessEditFormBytePlusConfigModelType = Partial<BytePlusAccessConfig>;

export type AccessEditFormBytePlusConfigProps = {
  form: FormInstance;
  disabled?: boolean;
  loading?: boolean;
  model?: AccessEditFormBytePlusConfigModelType;
  onModelChange?: (model: AccessEditFormBytePlusConfigModelType) => void;
};

const initModel = () => {
  return {} as AccessEditFormBytePlusConfigModelType;
};

const AccessEditFormBytePlusConfig = ({ form, disabled, loading, model, onModelChange }: AccessEditFormBytePlusConfigProps) => {
  const { t } = useTranslation();

  const formSchema = z.object({
    accessKey: z
      .string()
      .trim()
      .min(1, t("access.form.byteplus_access_key.placeholder"))
      .max(64, t("common.errmsg.string_max", { max: 64 })),
    secretKey: z
      .string()
      .trim()
      .min(1, t("access.form.byteplus_secret_key.placeholder"))
      .max(64, t("common.errmsg.string_max", { max: 64 })),
  });
  const formRule = createSchemaFieldRule(formSchema);

  const [initialValues, setInitialValues] = useState<Partial<z.infer<typeof formSchema>>>(model ?? initModel());
  useEffect(() => {
    setInitialValues(model ?? initModel());
  }, [model]);

  const handleFormChange = (_: unknown, fields: AccessEditFormBytePlusConfigModelType) => {
    onModelChange?.(fields);
  };

  return (
    <Form form={form} disabled={loading || disabled} initialValues={initialValues} layout="vertical" name="configForm" onValuesChange={handleFormChange}>
      <Form.Item
        name="accessKey"
        label={t("access.form.byteplus_access_key.label")}
        rules={[formRule]}
        tooltip={<span dangerouslySetInnerHTML={{ __html: t("access.form.byteplus_access_key.tooltip") }}></span>}
      >
        <Input placeholder={t("access.form.byteplus_access_key.placeholder")} />
      </Form.Item>

      <Form.Item
        name="secretKey"
        label={t("access.form.byteplus_secret_key.label")}
        rules={[formRule]}
        tooltip={<span dangerouslySetInnerHTML={{ __html: t("access.form.byteplus_secret_key.tooltip") }}></span>}
      >
        <Input.Password placeholder={t("access.form.byteplus_secret_key.placeholder")} />
      </Form.Item>
    </Form>
  );
};

export default AccessEditFormBytePlusConfig;
