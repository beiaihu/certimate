import { useEffect, useState } from "react";
import { useTranslation } from "react-i18next";
import { Form, Input, type FormInstance } from "antd";
import { createSchemaFieldRule } from "antd-zod";
import { z } from "zod";

import { type VolcEngineAccessConfig } from "@/domain/access";

type AccessEditFormVolcEngineConfigModelType = Partial<VolcEngineAccessConfig>;

export type AccessEditFormVolcEngineConfigProps = {
  form: FormInstance;
  disabled?: boolean;
  loading?: boolean;
  model?: AccessEditFormVolcEngineConfigModelType;
  onModelChange?: (model: AccessEditFormVolcEngineConfigModelType) => void;
};

const initModel = () => {
  return {} as AccessEditFormVolcEngineConfigModelType;
};

const AccessEditFormVolcEngineConfig = ({ form, disabled, loading, model, onModelChange }: AccessEditFormVolcEngineConfigProps) => {
  const { t } = useTranslation();

  const formSchema = z.object({
    accessKeyId: z
      .string()
      .trim()
      .min(1, t("access.form.volcengine_access_key_id.placeholder"))
      .max(64, t("common.errmsg.string_max", { max: 64 })),
    secretAccessKey: z
      .string()
      .trim()
      .min(1, t("access.form.volcengine_secret_access_key.placeholder"))
      .max(64, t("common.errmsg.string_max", { max: 64 })),
  });
  const formRule = createSchemaFieldRule(formSchema);

  const [initialValues, setInitialValues] = useState<Partial<z.infer<typeof formSchema>>>(model ?? initModel());
  useEffect(() => {
    setInitialValues(model ?? initModel());
  }, [model]);

  const handleFormChange = (_: unknown, fields: AccessEditFormVolcEngineConfigModelType) => {
    onModelChange?.(fields);
  };

  return (
    <Form form={form} disabled={loading || disabled} initialValues={initialValues} layout="vertical" name="configForm" onValuesChange={handleFormChange}>
      <Form.Item
        name="accessKeyId"
        label={t("access.form.volcengine_access_key_id.label")}
        rules={[formRule]}
        tooltip={<span dangerouslySetInnerHTML={{ __html: t("access.form.volcengine_access_key_id.tooltip") }}></span>}
      >
        <Input placeholder={t("access.form.volcengine_access_key_id.placeholder")} />
      </Form.Item>

      <Form.Item
        name="secretAccessKey"
        label={t("access.form.volcengine_secret_access_key.label")}
        rules={[formRule]}
        tooltip={<span dangerouslySetInnerHTML={{ __html: t("access.form.volcengine_secret_access_key.tooltip") }}></span>}
      >
        <Input.Password placeholder={t("access.form.volcengine_secret_access_key.placeholder")} />
      </Form.Item>
    </Form>
  );
};

export default AccessEditFormVolcEngineConfig;
