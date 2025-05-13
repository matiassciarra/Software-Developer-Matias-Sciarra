export function filterByFieldContains<T>(data: T[], field: keyof T, value: string): T[] {
  if (!value) return data; // Si no hay valor, devuelve todos los datos
  return data.filter((item) => {
    const fieldValue = item[field];
    return (
      typeof fieldValue === 'string' &&
      fieldValue.toLowerCase().includes(value.toLowerCase())
    );
  });
}