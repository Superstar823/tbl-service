import React, { useEffect, useState } from "react";
import { Button } from "antd";
import { HomePageContainer, SearchContainer } from "./styles";
import CustomTable from "../../components/Table";
import CustomModal from "../../components/Modal";
import { columns } from "../../types/constant";
import {
  fetchTransactions,
  createTransaction,
  updateTransaction,
  deleteTransaction,
} from "../../services/transactionApi";
import { fetchCustomers } from "../../services/customerApi";
import { fetchItems } from "../../services/itemApi";
import {
  iCustomerData,
  iItemData,
  iNewCustomerOrder,
  iTransactionData,
} from "../../types/types";

const HomePage: React.FC = () => {
  const [tblLoading, setTblLoading] = useState<boolean>(false);
  const [customerData, setCustomerData] = useState<iCustomerData[]>([]);
  const [itemData, setItemData] = useState<iItemData[]>([]);
  const [transactionData, setTransactionData] = useState<iTransactionData[]>(
    []
  );
  const [updateRowId, setUpdateRowId] = useState<string>("");
  const [isModalVisible, setIsModalVisible] = useState<boolean>(false);
  const [newCustomerOrder, setNewCustomerOrder] = useState<iNewCustomerOrder>({
    customer_id: "",
    item_id: "",
    price: 0,
    qty: 0,
    amount: 0,
  });

  useEffect(() => {
    fetchInitialData();
  }, []);

  const fetchInitialData = async () => {
    try {
      const [transactions, customers, items] = await Promise.all([
        fetchTransactions(),
        fetchCustomers(),
        fetchItems(),
      ]);
      setTransactionData(transactions);
      setCustomerData(customers);
      setItemData(items);
    } catch (error) {
      console.error("Error fetching initial data:", error);
    }
  };

  const showModal = () => {
    initialCustomerOrder();
    setIsModalVisible(true);
  };

  const initialCustomerOrder = () => {
    setUpdateRowId("");
    setNewCustomerOrder({
      customer_id: null,
      item_id: null,
      price: 0,
      qty: 0,
      amount: 0,
    });
  };

  const handleOk = async () => {
    const { customer_id, item_id, price, qty } = newCustomerOrder;
    setTblLoading(true);
    const newTransaction = {
      customer_id,
      item_id,
      price: Number(price),
      qty: Number(qty),
      amount: Number(price) * Number(qty),
    };
    try {
      if (updateRowId === "") await createTransaction(newTransaction);
      else await updateTransaction(updateRowId, newTransaction);

      await fetchInitialData();
      setTblLoading(false);
      setIsModalVisible(false);
    } catch (error) {
      console.error("Error creating new customer order:", error);
    }
  };

  const handleCancel = () => {
    setIsModalVisible(false);
  };

  const handleNewCustomerOrder = (name: string, value: string) => {
    setNewCustomerOrder((prev: iNewCustomerOrder) => ({
      ...prev,
      [name]: value,
    }));
  };

  const updateCustomerOrder = (record: iTransactionData) => {
    const { customer_id, item_id, price, qty, id } = record;
    setUpdateRowId(id);
    setNewCustomerOrder({
      customer_id,
      item_id,
      price,
      qty,
      amount: price * qty,
    });
    setIsModalVisible(true);
  };

  const deleteCustomerOrder = async (customerId: string) => {
    setTblLoading(true);
    try {
      await deleteTransaction(customerId);
      await fetchInitialData();
      await setTblLoading(false);
    } catch (error) {
      console.error("Error deleting customer order:", error);
    }
  };

  return (
    <HomePageContainer>
      <SearchContainer>
        <Button type="primary" className="success-button" onClick={showModal}>
          New
        </Button>
      </SearchContainer>
      <CustomTable
        columns={columns}
        data={transactionData}
        loading={tblLoading}
        setUpdateRowId={setUpdateRowId}
        deleteCustomerOrder={deleteCustomerOrder}
        updateCustomerOrder={updateCustomerOrder}
      />
      <CustomModal
        isVisible={isModalVisible}
        customerData={customerData}
        itemData={itemData}
        handleOk={handleOk}
        handleCancel={handleCancel}
        newCustomerOrder={newCustomerOrder}
        handleNewCustomerOrder={handleNewCustomerOrder}
      />
    </HomePageContainer>
  );
};

export default HomePage;
