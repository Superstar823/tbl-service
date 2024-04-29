// components/Table.tsx
import React from "react";
import { Table, Button, Popconfirm } from "antd";
import { OptionContainer } from "./styles";
import { iCustomTable, iTransactionData } from "../../types/types";
import formatDateTime from "../../utils/formatDateTime";

const renderIdColumn = (_: iTransactionData, index: number) => (
  <span>{index + 1}</span>
);

const renderOperationsColumn = (
  record: iTransactionData,
  updateCustomerOrder: Function,
  deleteCustomerOrder: Function
) => (
  <OptionContainer>
    <Button
      type="primary"
      size="small"
      onClick={() => updateCustomerOrder(record)}
    >
      Edit
    </Button>
    <Popconfirm
      title="Delete the transaction"
      placement="bottomRight"
      description="Are you sure to delete this transaction?"
      okText="Yes"
      cancelText="No"
      onConfirm={() => deleteCustomerOrder(record.id)}
    >
      <Button type="primary" className="danger-button" size="small">
        Delete
      </Button>
    </Popconfirm>
  </OptionContainer>
);

const renderPriceColumn = (value: number) => <span>{value}$</span>;

const renderTotalColumn = (record: iTransactionData) => (
  <span>{record.price * record.qty}$</span>
);

const renderTimeColumn = (value: string) => (
  <span>{formatDateTime(value)}</span>
);

const renderColumn = (
  columnKey: string,
  columns: any[],
  updateCustomerOrder: Function,
  deleteCustomerOrder: Function
) => {
  const column = columns.find((col) => col.key === columnKey);
  if (!column) return null;

  switch (columnKey) {
    case "id":
      column.render = (_: string, record: iTransactionData, index: number) =>
        renderIdColumn(record, index);
      break;
    case "operations":
      column.render = (_: string, record: iTransactionData) =>
        renderOperationsColumn(
          record,
          updateCustomerOrder,
          deleteCustomerOrder
        );
      break;
    case "price":
      column.render = (value: number) => renderPriceColumn(value);
      break;
    case "total":
      column.render = (_: string, record: iTransactionData) =>
        renderTotalColumn(record);
      break;
    case "createTime":
    case "modifyTime":
      column.render = (value: string) => renderTimeColumn(value);
      break;
    default:
      break;
  }
};

const CustomTable: React.FC<iCustomTable> = ({
  columns,
  data,
  loading,
  updateCustomerOrder,
  deleteCustomerOrder,
}) => {
  const newData =
    data &&
    data.map((record: iTransactionData, index: number) => ({
      ...record,
      primaryId: index + 1
    }));

  columns.forEach((column) =>
    renderColumn(column.key, columns, updateCustomerOrder, deleteCustomerOrder)
  );

  return (
    <Table
      columns={columns}
      loading={loading}
      dataSource={newData}
      bordered
      size="small"
      pagination={false}
      rowKey="id"
    />
  );
};

export default CustomTable;
