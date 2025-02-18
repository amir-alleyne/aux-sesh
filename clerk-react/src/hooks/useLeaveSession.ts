import { useCallback } from "react";
import {  leaveSession } from "../api/sessions";
import { useApi } from "./useApi";

export function useLeaveSession() {
    // The generic hook infers the types for data and arguments.
    const { data, loading, error, callApi } = useApi<string, [Number, string]>(leaveSession);
  
    // Optionally, wrap callApi to create a clearer API for the hook user.
    const leaveAuxSession = useCallback((session_id: Number, token: string) => callApi(session_id, token), [callApi]);
  
  
    return { leaveResponse: data, loading, error, leaveAuxSession };
  }