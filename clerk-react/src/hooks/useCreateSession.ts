import { useApi } from './useApi';
import { createSession, CreateSessionResponse, SessionData } from '../api/sessions';

export function useCreateSession() {
  // The generic hook infers the types for data and arguments.
  const { data, loading, error, callApi } = useApi<CreateSessionResponse, []>(createSession);

  // Optionally, wrap callApi to create a clearer API for the hook user.
  const createAuxSession = () => callApi();

  return { session: data, loading, error, createAuxSession };
}