// components/Modal/index.tsx
import React, { useEffect, useMemo, useState } from "react";
import { Modal, Select, Input } from "antd";
import { ModalContentWrapper, Row } from "./styles";
import { iCustomModal, iCustomerData, iItemData } from "../../types/types";

const { Option } = Select;

const CustomModal: React.FC<iCustomModal> = ({
  isVisible,
  customerData,
  itemData,
  newCustomerOrder,
  handleOk,
  handleCancel,
  handleNewCustomerOrder,
}) => {
  const [errors, setErrors] = useState<{ [key: string]: string }>({});

  useEffect(() => {
    setErrors({});
  }, [isVisible]);

  const validateInputs = () => {
    const newErrors: { [key: string]: string } = {};

    if (!newCustomerOrder.customer_id) {
      newErrors.customer_id = "Please select a customer";
    }

    if (!newCustomerOrder.item_id) {
      newErrors.item_id = "Please select an item";
    }

    if (
      !newCustomerOrder.price ||
      isNaN(newCustomerOrder.price) ||
      Number(newCustomerOrder.price) <= 0
    ) {
      newErrors.price = "Price must be a valid positive number";
    }

    if (
      !newCustomerOrder.qty ||
      isNaN(newCustomerOrder.qty) ||
      Number(newCustomerOrder.qty) <= 0
    ) {
      newErrors.qty = "Quantity must be a valid positive number";
    }

    setErrors(newErrors);
    return Object.keys(newErrors).length === 0;
  };

  useEffect(() => {
    const originalConsoleError = console.error;
    console.error = (...args: any[]) => {
      if (
        typeof args[0] === "string" &&
        args[0].includes("findDOMNode is deprecated")
      ) {
        return;
      }
      originalConsoleError.call(console, ...args);
    };

    return () => {
      console.error = originalConsoleError;
    };
  }, []);

  const handleInputChange = (
    key: keyof typeof newCustomerOrder,
    value: string
  ) => {
    handleNewCustomerOrder(key, value);
  };

  const handleKeyDown = (e: React.KeyboardEvent<HTMLInputElement>) => {
    if (!/\d/.test(e.key) && e.key !== "." && e.key !== "Backspace") {
      e.preventDefault();
    }
  };

  const handleOkClick = () => {
    if (validateInputs()) {
      handleOk();
    } else {
      console.log("Please fix the errors before saving.");
    }
  };

  const total = useMemo(
    () => newCustomerOrder.price * newCustomerOrder.qty,
    [newCustomerOrder.price, newCustomerOrder.qty]
  );

  return (
    <Modal
      title="Customer Order"
      open={isVisible}
      okText="Save"
      onOk={handleOkClick}
      onCancel={handleCancel}
    >
      <ModalContentWrapper>
        <Row>
          <label className="label mt-4">Customer</label>
          <Select
            style={{ width: "100%" }}
            status={errors.customer_id ? "error" : ""}
            placeholder="Select Customer Name"
            onChange={(customerID: string) =>
              handleInputChange("customer_id", customerID)
            }
            value={newCustomerOrder.customer_id}
          >
            {customerData?.map(({ id, customer_name }: iCustomerData) => (
              <Option value={id} key={id}>
                {customer_name}
              </Option>
            ))}
          </Select>
        </Row>
        <Row>
          <label className="label mt-4">Item Name</label>
          <Select
            style={{ width: "100%" }}
            status={errors.item_id && "error"}
            placeholder="Select Item Name"
            onChange={(itemID: string) => handleInputChange("item_id", itemID)}
            value={newCustomerOrder.item_id}
          >
            {itemData?.map(({ id, item_name }: iItemData) => (
              <Option value={id} key={id}>
                {item_name}
              </Option>
            ))}
          </Select>
        </Row>
        <Row>
          <label className="label mt-4">Price</label>
          <Input
            type="number"
            status={errors.price && "error"}
            placeholder="Please enter the Price"
            value={newCustomerOrder.price}
            min={0}
            onKeyDown={handleKeyDown}
            onChange={(e) => handleInputChange("price", e.target.value)}
          />
        </Row>
        <Row>
          <label className="label mt-4">QTY</label>
          <Input
            placeholder="Please enter the QTY"
            status={errors.qty && "error"}
            value={newCustomerOrder.qty}
            min={0}
            onKeyDown={handleKeyDown}
            onChange={(e) => handleInputChange("qty", e.target.value)}
          />
        </Row>
        <Row>
          <label className="label">Total:</label>
          <span>
            <b>{total}$</b>
          </span>
        </Row>
      </ModalContentWrapper>
    </Modal>
  );
};

export default CustomModal;
