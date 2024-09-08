import { useState } from 'react';
import { Button, Modal, Form, Input, Space } from 'antd';
import { ProFormText } from '@ant-design/pro-components';
import { EditOutlined, MinusCircleOutlined, PlusOutlined } from '@ant-design/icons';
import './style.scss'
import http from '../../../config';
import { toast } from 'react-toastify';

function Index(props: any) {
  const [isModalOpen, setIsModalOpen] = useState(false);
  const [form] = Form.useForm();

  const showModal = () => {
    form.resetFields();
    if (props?.data) {
      form.setFieldsValue({
        title: props.data.title,
        options: props.data.options.map((option: any) => ({
          variant: option.variant,
          ball: option.ball.toString(), // Convert ball to string if it's a number
        })),
      });
    }
    setIsModalOpen(true);
  };

  const handleOk = () => {
    setIsModalOpen(false);
  };

  const handleCancel = () => {
    setIsModalOpen(false);
  };

  const onFinish = async (values: any) => {
    const formattedValues = {
      ...values,
      options: values.options.map((option: any) => ({
        ...option,
        ball: Number(option.ball),
      })),
    };
    formattedValues.id = props?.data.id
    try{
      const response = await http.put('/poll/', formattedValues)
      if(response.status == 200){
        toast.success("To'plam tahrirlandi!");
      }
    }catch(err){
      toast.error("To'plam tahrirlashda qandaydir muommo yuzaga keldi")
      console.error(err)
    }
    props?.getData();
    setIsModalOpen(false);
    form.resetFields();
  };

  const getAlphabetLabel = (index: number) => {
    return String.fromCharCode(65 + index);
  };

  return (
    <>
      <EditOutlined style={{ position: 'absolute', top: 13, right: 40 }} onClick={showModal} />
      <Modal title={"So'rovnoma to'plam tahrirlash"} open={isModalOpen} onOk={handleOk} onCancel={handleCancel} footer={null}>
        <Form
          form={form}
          onFinish={onFinish}
          initialValues={{
            options: [{ variant: '', ball: '' }]
          }}
        >
          <ProFormText
            initialValue={''}
            hasFeedback
            name="title"
            placeholder="Iltimos so'rovnoma to'plami nomini kiriting"
            rules={[
              {
                required: true,
                message: "To'plam nomini kiriting!",
              },
            ]}
          />

          <Form.List name="options">
            {(fields, { add, remove }) => (
              <>
                {fields.map(({ key, name, ...restField }, index) => (
                  <Space key={key} style={{ display: 'block', marginBottom: 8, width: '100%', position: 'relative' }} align="baseline">
                    <span style={{ position: 'absolute', top: 5, left: 0, fontWeight: 700 }}>
                      {getAlphabetLabel(index)}
                    </span>
                    <div style={{ display: 'flex-box', alignItems: 'center', width: '100%', position: 'relative' }}>
                      <Form.Item
                        {...restField}
                        name={[name, 'variant']}
                        rules={[{ required: true, message: 'Variant kiriting!' }]}
                        style={{ flex: 1, width: '70%', paddingLeft: 20 }}
                      >
                        <Input style={{ width: '100%' }} placeholder="Variant kiriting (masalan: har doim, odatda)" />
                      </Form.Item>
                      <Form.Item
                        {...restField}
                        name={[name, 'ball']}
                        rules={[{ required: true, message: 'Ball kiriting!' }]}
                        style={{ flex: 1, width: '28%', position: 'absolute', right: 0, top: 0 }}
                      >
                        <Input style={{ width: '100%' }} placeholder="Ball kiritng" />
                      </Form.Item>
                    </div>
                    {fields.length > 1 && (
                      <MinusCircleOutlined style={{ position: 'absolute', top: 10, right: 10 }} onClick={() => remove(name)} />
                    )}
                  </Space>
                ))}
                <Form.Item>
                  <Button type="dashed" onClick={() => add()} block icon={<PlusOutlined />}>
                    Variant qo'shish
                  </Button>
                </Form.Item>
              </>
            )}
          </Form.List>

          <div style={{ display: 'flex', justifyContent: 'end' }}>
            <Button type="primary" htmlType="submit">
              Yaratish
            </Button>
          </div>
        </Form>
      </Modal>
    </>
  );
}

export default Index;
