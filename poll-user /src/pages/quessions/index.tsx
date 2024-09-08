import { MenuIds } from "@store";
import { useEffect, useState } from "react";
import http from "../../config";
import { Button, Form, Modal, Radio, Spin } from "antd";
import './style.scss';
import { LoadingOutlined } from "@ant-design/icons";
import { toast, ToastContainer } from "react-toastify";

function Index() {
  const { menu_id }:any = MenuIds(); 
  const [data, setData] = useState<any[]>([]); 
  const [variant, setVariant] = useState<any[]>([]);
  const [answers, setAnswers] = useState<{ question_id: string; answer_point: number }[]>([]);
  const [poll, setPoll] = useState<any>({});
  const [load, setLoad] = useState(false);
  const [feedback, setFeedback] = useState("")
  
  
  const [isModalOpen, setIsModalOpen] = useState(false);

  const showModal = () => {
    setIsModalOpen(true);
  };

  const handleOk = () => {
    setIsModalOpen(false);
  };

  const handleCancel = () => {
    setIsModalOpen(false);
  };
  
  const [form] = Form.useForm(); 

  async function getData() {
    try {
      if (menu_id) {
        const response = await http.get(`/questions/${menu_id}`);
        setPoll(response?.data?.poll || {}); 
        setData(response?.data?.question || []);
        setVariant(response?.data?.poll?.options || []); 
      }
    } catch (e) {
      console.error('Error fetching data:', e);
    }
  }

  useEffect(() => {
    getData();
  }, [menu_id]);

  const getLetter = (index: number) => String.fromCharCode(65 + index);


  const handleAnswerChange = (question_id: string, answer_point: number) => {
    setAnswers(prevAnswers => {
      const existingAnswerIndex = prevAnswers.findIndex(answer => answer.question_id === question_id);
      if (existingAnswerIndex !== -1) {
        const updatedAnswers = [...prevAnswers];
        updatedAnswers[existingAnswerIndex] = { question_id, answer_point };
        return updatedAnswers;
      }
      return [...prevAnswers, { question_id, answer_point }];
    });
  };

  const handleSubmit = async (_: any) => {
    const payload = {
      answers: answers,
      poll_id: poll?.id,
      user_id: localStorage.getItem('user_id'),
    };
    try {
      const response = await http.post('/send-answers', payload);
      if (response.status === 201) {
        setFeedback(response?.data?.feedback)
        showModal()
        form.resetFields(); 
      }
    } catch (err) {
      toast.error("Ushbu sizning testingizni tasdiqlashda xatolik yuz berdi!", { autoClose: 1500 });
      console.error('Error submitting answers:', err);
    }
  };

  useEffect(() => {
    if (menu_id) {
      setLoad(true);
      setTimeout(() => setLoad(false), 700);
    }
  }, [menu_id]);

  return (
    load ? (
      <div style={{ display: 'flex', justifyContent: 'center', alignItems: 'center', marginTop: 150 }}>
        <Spin indicator={<LoadingOutlined style={{ fontSize: 48 }} spin />} />
      </div>
    ) : (
      <>
        <ToastContainer />
        <div className="question">
          <h2 style={{ textAlign: 'center', marginBottom: 40 }}>{poll?.title}</h2>
          {data.length > 0 ? (
            <Form form={form} onFinish={handleSubmit}>
              {data.map((e, i) => (
                <div key={i} className="test">
                  <p>{i + 1}. {e.content}</p>
                  <Form.Item
                    name={`answer${i}`}
                    rules={[{
                      required: true,
                      message: "Iltimos savollarni to'ldiring!",
                    }]}
                  >
                    <Radio.Group
                      className="radio"
                      size="large"
                      onChange={(event) => handleAnswerChange(e.id, event.target.value)}
                    >
                      {variant.map((option, index) => (
                        <Radio key={index} value={option?.ball}>
                          {`${getLetter(index)}. ${option?.variant}`}
                        </Radio>
                      ))}
                    </Radio.Group>
                  </Form.Item>
                </div>
              ))}
              {menu_id && (
                <Button
                  style={{ display: 'block', width: '100%', maxWidth: 100, margin: '0 auto' }}
                  type="primary"
                  htmlType="submit"
                >
                  Tasdiqlash
                </Button>
              )}
            </Form>
          ) : (
            <p style={{textAlign: 'center', fontSize: 20, fontWeight: 600}}>Savollar mavjud emas !</p>
          )}
        </div>

        <Modal title="Javob" open={isModalOpen} onOk={handleOk} footer={false} onCancel={handleCancel}>
            <p>{feedback}</p>
        </Modal>
      </>
    )
  );
}

export default Index;
