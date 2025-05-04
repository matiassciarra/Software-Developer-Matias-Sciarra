import axios from 'axios';

const apiClient = axios.create({
  baseURL: import.meta.env.VITE_URL_API_BACKEND,
  timeout: 5000,
});

export interface Item {
  ID: bigint;             // Identificador único
  Action: string;         // Acción realizada (e.g., "reiterated by")
  Brokerage: string;      // Nombre de la correduría (e.g., "Benchmark")
  Company: string;        // Nombre de la compañía (e.g., "Hello Group")
  RatingFrom: string;     // Calificación original (e.g., "Buy")
  RatingTo: string;       // Calificación actualizada (e.g., "Hold")
  TargetFrom: number;     // Precio objetivo anterior
  TargetTo: number;       // Precio objetivo actualizado
  Ticker: string;         // Código de bolsa (e.g., "MOMO")
  Time: string;           // Fecha y hora en formato ISO 8601
}

// Función para obtener los datos (GET request)
export async function fetchItems(): Promise<Item[]> {
  try {
    const response = await apiClient.get<Item[]>('/items'); // Llama al endpoint /items
    return response.data; // Retorna los datos
  } catch (error) {
    console.error('Error al obtener los datos:', error);
    throw error; // Lanza el error para manejarlo en el componente
  }
}