import { useCallback } from "react";
import { JoinSessionResponse, joinSession } from "../api/sessions";
import { useApi } from "./useApi";

export function useJoinSession() {
    // The generic hook infers the types for data and arguments.
    const { data, loading, error, callApi } = useApi<JoinSessionResponse, [Number, string]>(joinSession);
  
    // Optionally, wrap callApi to create a clearer API for the hook user.
    const joinAuxSession = useCallback((session_id: Number, token: string) => callApi(session_id, token), [callApi]);
  
  
    return { joinResponse: data, loading, error, joinAuxSession };
  }