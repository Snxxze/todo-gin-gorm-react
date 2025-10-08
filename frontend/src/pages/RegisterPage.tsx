import React, { useState } from "react";
import {
  Form,
  Input,
  Button,
  Card, 
  Typography,
  message,
} from "antd";
import { useNavigate } from "react-router-dom";
import { register } from "../services/auth"
import type { RegisterRequest } from "../interfaces/Auth";

const { Title } = Typography;

const RegisterPage: React.FC = () => {
  const [loading, setLoading] = useState(false);
  const navigate = useNavigate();

  const onFinish = async (values: RegisterRequest) => {
    setLoading(true);
    try {
      const data = await register(values)
      message.success("Registration successful!")
      navigate("/login")
    } catch (err: any) {
      message.error(err.response?.data?.error || "Registration failed")
    } finally {
      setLoading(false)
    }
  }

  return (
    <div
      style={{
        height: "100vh",
        display: "flex",
        alignItems: "center",
        justifyContent: "center",
        background: "#f5f5f5",
      }}
    >
      <Card
        style={{
          width: 350,
          boxShadow: "0 2px 10px rgba(0,0,0,0.1)"
        }}
      >
        <Title
          level={3}
          style={{
            textAlign: "center"
          }}
        >
          Create Account
        </Title>
        
        <Form layout="vertical" onFinish={onFinish}>
          <Form.Item
            name="email"
            label="Email"
            rules={[{ required: true, type: "email"}]}
          >
            <Input placeholder="Enter your email"/>
          </Form.Item>

          <Form.Item
            name="password"
            label="Password"
            rules={[{ required: true, min: 6 }]}
          >
            <Input placeholder="Enter your password"/>
          </Form.Item>

          <Button
            type="primary"
            htmlType="submit"
            block
            loading={loading}
            style={{marginTop: "10px"}}
          >
            Register
          </Button>
        </Form>

        <div
          style={{
            marginTop: "10px", 
            textAlign: "center"
          }}
        >
          <a onClick={() => navigate("/login")}>Already have an account?</a>
        </div>
      </Card>
    </div>
  )
}

export default RegisterPage