import { Table, Radio } from 'antd';
import './style.scss';

function Index({ data, columns, options }: any) {
  const dataWithKey = data.map((item: any, index: number) => ({
    ...item,
    key: item.id || index,
  }));

  return (
    <>
      <Table
        columns={columns}
        dataSource={dataWithKey}
        pagination={false}
        expandable={{
          expandedRowRender: (_, i) => {
            return (
              <Radio.Group key={i + 1}>
                {options
                  ?.sort((a: any, b: any) => a.variant.localeCompare(b.variant))
                  ?.map((opt: any, index: any) => (
                    <Radio key={index} value={opt.ball}>
                      {String.fromCharCode(65 + index)}. {opt.variant}
                    </Radio>
                  ))}
              </Radio.Group>
            );
          },
        }}
      />
    </>
  );
}

export default Index;
