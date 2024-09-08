import { useState } from 'react';
import { Button, Modal, Form } from 'antd';
import { ProFormText } from '@ant-design/pro-components';
import './styls.scss'
import http from '../../../config';
import { MenuIds } from '@store';
import { toast } from 'react-toastify';

function Index(props:any) {
  const [isModalOpen, setIsModalOpen] = useState(false);
  const [form] = Form.useForm();
  const {menu_id}:any = MenuIds()

  const showModal = () => {
    form.resetFields();
    setIsModalOpen(true);
  };


  const handleCancel = () => {
    setIsModalOpen(false);
  };

  async function handleSubmit(values: any){
    if(props?.title){
      try{
        values.id = props?.data?.id
        const response = await http.put(`/question/`, values)
        if(response.status == 200){
          toast.success("Savol muvaffqiyatli tahrirlandi", {autoClose: 1200})
        }
      }catch(err){
          toast.error("Savol tahrirlashda qandaydir muommo yuzaga keldi", {autoClose: 1200})
      }
    }else{
      try{
        values.poll_id = menu_id
        const response = await http.post('/question/', values)
        if(response.status == 200){
          toast.success("Savol muvaffqiyatli qo'shildi", {autoClose: 1200})
        }
      }catch(err){
          toast.error("Savol qo'shishda qandaydir muommo yuzaga keldi", {autoClose: 1200})
      }
    }
    props?.getData()
    handleCancel();
  }

  const style = {
    marginBottom: props?.title ? 0 : 20,
  }

  return (
    <>
      <Button type={props?.title ? 'dashed' : 'primary'} style={{ width: '100%', maxWidth: '200px', ...style }} onClick={showModal}>
        {props?.title || "Savol qo'shish"} 
      </Button>
      <Modal
        title={props?.modaltitle ? 'Savol Taxrirlash' : "Savol qo'shish"}
        open={isModalOpen}
        onCancel={handleCancel}
        footer={null}
      >
        <Form onFinish={(values) => handleSubmit(values)} form={form} layout="vertical">
          <ProFormText
            initialValue={props?.data?.content || ''}
            name="content"
            label="Savol nomi"
            placeholder="Iltimos so'rovnoma nomini kiriting"
            rules={[{ required: true, message: "So'rovnoma nomini kiriting !" }]}
          />

          <div style={{ display: 'flex', justifyContent: 'end' }}>
            <Button htmlType='submit' type="primary">
              Yaratish
            </Button>
          </div>
        </Form>
      </Modal>
    </>
  );
}

export default Index;
