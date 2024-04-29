export type iColumn = {
  title: string;
  dataIndex: string;
  key: string;
};

export type iCustomModal = {
  isVisible: boolean;
  customerData: iCustomerData[];
  itemData: iItemData[];
  newCustomerOrder: iNewCustomerOrder;
  handleOk: () => void;
  handleCancel: () => void;
  handleNewCustomerOrder: (fieldName: string, value: string) => void;
};

export type iCustomTable = {
  columns: any[];
  data: iTransactionData[];
  loading: boolean;
  updateCustomerOrder: (record: iTransactionData) => void;
  deleteCustomerOrder: (customerId: string) => void;
  setUpdateRowId: (rowId: string) => void;
};

export type iNewCustomerOrder = {
  customer_id: string | null;
  item_id: string | null;
  price: number;
  qty: number;
  amount: number;
};

export type iItemData = {
  id: string;
  item_name: string;
  cost: number;
  price: number;
  sort: number;
  created_at: string;
  updated_at: string;
};

export type iCustomerData = {
  id: string;
  customer_name: string;
  balance: number;
  created_at: string;
  updated_at: string;
};

export type iTransactionData = {
  primaryId?: number;
  id: string;
  customer_id: string;
  customer_name: string;
  item_id: string;
  item_name: string;
  qty: number;
  price: number;
  amount: number;
  created_at: string;
  updated_at: string;
};
