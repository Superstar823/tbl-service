// api/transactionApi.ts
import axios from "axios";
import { BASE_URL } from "../types/constant";
import { iNewCustomerOrder, iTransactionData } from "../types/types";

const transactionApi = axios.create({
  baseURL: `${BASE_URL}/transactions`
});

export const fetchTransactions = async (): Promise<iTransactionData[]> => {
  try {
    const response = await transactionApi.get<iTransactionData[]>("");
    return response.data;
  } catch (error) {
    throw new Error(`Error fetching transactions: ${error}`);
  }
};

export const createTransaction = async (newCustomerOrderData: Partial<iNewCustomerOrder>): Promise<iNewCustomerOrder> => {
  try {
    const response = await transactionApi.post<iNewCustomerOrder>("", newCustomerOrderData);
    return response.data;
  } catch (error) {
    throw new Error(`Error creating transaction: ${error}`);
  }
};

export const updateTransaction = async (rowId: string, updateCustomerOrderData: Partial<iNewCustomerOrder>): Promise<iNewCustomerOrder> => {
  try {
    const updateData = {...updateCustomerOrderData, id:rowId}
    const response = await transactionApi.put<iNewCustomerOrder>(``, updateData);
    return response.data;
  } catch (error) {
    throw new Error(`Error updating transaction: ${error}`);
  }
};

export const deleteTransaction = async (id: string): Promise<void> => {
  try {
    await transactionApi.delete(`/${id}`);
  } catch (error) {
    throw new Error(`Error deleting transaction: ${error}`);
  }
};