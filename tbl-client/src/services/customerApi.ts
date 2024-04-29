// api/customerApi.ts
import axios from "axios";
import { BASE_URL } from "../types/constant";
import { iCustomerData } from "../types/types";

const customerApi = axios.create({
  baseURL: `${BASE_URL}/customers`
});

export const fetchCustomers = async (): Promise<iCustomerData[]> => {
  try {
    const response = await customerApi.get<iCustomerData[]>("");
    return response.data;
  } catch (error) {
    throw new Error(`Error fetching customers: ${error}`);
  }
};

export const createCustomer = async (newCustomerOrderData: Partial<iCustomerData>): Promise<iCustomerData> => {
  try {
    const response = await customerApi.post<iCustomerData>("", newCustomerOrderData);
    return response.data;
  } catch (error) {
    throw new Error(`Error creating customer: ${error}`);
  }
};

export const updateCustomer = async (rowId: string, updateCustomerOrderData: Partial<iCustomerData>): Promise<iCustomerData> => {
  try {
    const updateData = {...updateCustomerOrderData, id:rowId}
    const response = await customerApi.put<iCustomerData>(``, updateData);
    return response.data;
  } catch (error) {
    throw new Error(`Error updating customer: ${error}`);
  }
};

export const deleteCustomer = async (id: string): Promise<void> => {
  try {
    await customerApi.delete(`/${id}`);
  } catch (error) {
    throw new Error(`Error deleting customer: ${error}`);
  }
};