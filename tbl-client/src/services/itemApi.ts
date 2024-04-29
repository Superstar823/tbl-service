// api/itemApi.ts
import axios from "axios";
import { BASE_URL } from "../types/constant";
import { iItemData } from "../types/types";

const itemApi = axios.create({
  baseURL: `${BASE_URL}/items`
});

export const fetchItems = async (): Promise<iItemData[]> => {
  try {
    const response = await itemApi.get<iItemData[]>("");
    return response.data;
  } catch (error) {
    throw new Error(`Error fetching items: ${error}`);
  }
};

export const createItem = async (newItemData: Partial<iItemData>): Promise<iItemData> => {
  try {
    const response = await itemApi.post<iItemData>("", newItemData);
    return response.data;
  } catch (error) {
    throw new Error(`Error creating item: ${error}`);
  }
};

export const updateItem = async (rowId: string, updateItemData: Partial<iItemData>): Promise<iItemData> => {
  try {
    const updateData = {...updateItemData, id:rowId}
    const response = await itemApi.put<iItemData>(``, updateData);
    return response.data;
  } catch (error) {
    throw new Error(`Error updating item: ${error}`);
  }
};

export const deleteItem = async (id: string): Promise<void> => {
  try {
    await itemApi.delete(`/${id}`);
  } catch (error) {
    throw new Error(`Error deleting item: ${error}`);
  }
};