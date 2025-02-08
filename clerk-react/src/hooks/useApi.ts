import { useState, useCallback } from 'react';

export interface ApiState<T> {
  data?: T;
  loading: boolean;
  error?: Error;
}

export const BACKEND_BASE_URL = 'http://localhost:8080';

export function useApi<T, Args extends any[]>(apiFunction: (...args: Args) => Promise<T>) {
  const [data, setData] = useState<T | undefined>(undefined);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<Error | undefined>(undefined);

  const callApi = useCallback(async (...args: Args): Promise<T> => {
    setLoading(true);
    setError(undefined);
    try {
      const result = await apiFunction(...args);
      setData(result);
      return result;
    } catch (err) {
      setError(err as Error);
      throw err;
    } finally {
      setLoading(false);
    }
  }, [apiFunction]);

  return { data, loading, error, callApi };
}