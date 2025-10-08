import React, { useEffect, useState } from "react";
import {
  List,
  Button,
  Modal,
  Form,
  Input,
  Tag,
  message,
  Typography,
  Space,
} from "antd";
import {
  get,
  create,
  update,
  deleteTodo,
} from "../services/todo";
import type { Todo, CreateTodoRequest, UpdateTodoRequest } from "../interfaces/Todo";
import { logout } from "../services/auth";
import { useNavigate } from "react-router-dom";

const { Title } = Typography;

const TodoPage: React.FC = () => {
  const [todos, setTodos] = useState<Todo[]>([]);
  const [loading, setLoading] = useState(false);
  const [modalOpen, setModalOpen] = useState(false);
  const [editingTodo, setEditingTodo] = useState<Todo | null>(null);
  const [form] = Form.useForm();
  const navigate = useNavigate();

  // โหลดรายการ todo
  const loadTodos = async () => {
    setLoading(true);
    try {
      const data = await get();
      setTodos(data);
    } catch (err: any) {
      message.error(err.response?.data?.error || "Failed to load todos");
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    loadTodos();
  }, []);

  // เพิ่มหรือแก้ไข todo
  const handleSubmit = async (values: CreateTodoRequest | UpdateTodoRequest) => {
    try {
      if (editingTodo) {
        await update(editingTodo.id, values);
        message.success("Todo updated");
      } else {
        await create(values as CreateTodoRequest);
        message.success("Todo created");
      }
      console.log(values);
      setModalOpen(false);
      setEditingTodo(null);
      form.resetFields();
      loadTodos();
    } catch (err: any) {
      message.error(err.response?.data?.error || "Operation failed");
    }
  };

  // ลบ todo
  const handleDelete = async (id: number) => {
    Modal.confirm({
      title: "Are you sure you want to delete?",
      okText: "Yes",
      okType: "danger",
      cancelText: "No",
      onOk: async () => {
        try {
          await deleteTodo(id);
          message.success("Todo deleted");
          loadTodos();
        } catch (err: any) {
          message.error(err.response?.data?.error || "Delete failed");
        }
      },
    });
  };

  // เปลี่ยนสถานะ (pending ↔ done)
  const toggleStatus = async (todo: Todo) => {
    try {
      await update(todo.id, {
        status: todo.status === "done" ? "pending" : "done",
      });
      loadTodos();
    } catch (err: any) {
      message.error("Failed to update status");
    }
  };

  // ออกจากระบบ
  const handleLogout = () => {
    logout();
    message.info("Logged out");
    navigate("/login");
  };

  return (
    <div style={{ padding: 30, maxWidth: 700, margin: "0 auto" }}>
      {/* Header */}
      <Space
        style={{
          display: "flex",
          justifyContent: "space-between",
          marginBottom: 20,
        }}
      >
        <Title level={2}>My Todos</Title>
        <Button danger onClick={handleLogout}>
          Logout
        </Button>
      </Space>

      {/* ปุ่มเพิ่มงาน */}
      <Button
        type="primary"
        onClick={() => {
          setEditingTodo(null);
          setModalOpen(true);
          form.resetFields();
        }}
        style={{ marginBottom: 20 }}
      >
        + Add Todo
      </Button>

      {/* รายการ Todo */}
      <List
        bordered
        loading={loading}
        dataSource={todos}
        renderItem={(item) => (
          <List.Item
            actions={[
              <Button
                type="link"
                onClick={() => toggleStatus(item)}
                style={{ padding: 0 }}
              >
                {item.status === "done" ? "Mark Pending" : "Mark Done"}
              </Button>,
              <Button
                type="link"
                onClick={() => {
                  setEditingTodo(item);
                  form.setFieldsValue(item);
                  setModalOpen(true);
                }}
                style={{ padding: 0 }}
              >
                Edit
              </Button>,
              <Button
                danger
                type="link"
                onClick={() => handleDelete(item.id)}
                style={{ padding: 0 }}
              >
                Delete
              </Button>,
            ]}
          >
            <List.Item.Meta
              title={
                <span
                  style={{
                    textDecoration:
                      item.status === "done" ? "line-through" : "none",
                    color: item.status === "done" ? "#888" : "#000",
                  }}
                >
                  {item.title}
                </span>
              }
              description={item.description}
            />
            <Tag color={item.status === "done" ? "green" : "orange"}>
              {item.status.toUpperCase()}
            </Tag>
          </List.Item>
        )}
      />

      {/* Modal เพิ่ม / แก้ไข */}
      <Modal
        title={editingTodo ? "Edit Todo" : "Add Todo"}
        open={modalOpen}
        onCancel={() => setModalOpen(false)}
        footer={null}
      >
        <Form
          form={form}
          layout="vertical"
          onFinish={handleSubmit}
          initialValues={{ status: "pending" }}
        >
          <Form.Item
            name="title"
            label="Title"
            rules={[{ required: true, message: "Please enter a title" }]}
          >
            <Input placeholder="Enter todo title" />
          </Form.Item>

          <Form.Item name="description" label="Description">
            <Input.TextArea placeholder="Enter description" />
          </Form.Item>

          <Button type="primary" htmlType="submit" block>
            {editingTodo ? "Save Changes" : "Add Todo"}
          </Button>
        </Form>
      </Modal>
    </div>
  );
};

export default TodoPage;
