/**
 * Authentication hook for the Lilo outfit finder application
 */
import { useEffect, useState } from "react";
import api from "../../services/api";
import { User } from "../../types";

interface AuthState {
  user: User | null;
  isLoading: boolean;
  isAuthenticated: boolean;
  error: string | null;
}

export function useAuth() {
  const [state, setState] = useState<AuthState>({
    user: null,
    isLoading: true,
    isAuthenticated: false,
    error: null,
  });

  const checkAuthStatus = async () => {
    try {
      console.log("Checking auth status...");
      const response = await api.auth.getUser();
      console.log("Auth status response:", response);

      if (response.status === 200 && response.data) {
        console.log("Setting authenticated state with user:", response.data);
        setState({
          user: response.data,
          isLoading: false,
          isAuthenticated: true,
          error: null,
        });
      } else {
        console.log("Setting unauthenticated state");
        setState({
          user: null,
          isLoading: false,
          isAuthenticated: false,
          error: null,
        });
      }
    } catch (error) {
      console.error("Auth status check error:", error);
      setState({
        user: null,
        isLoading: false,
        isAuthenticated: false,
        error:
          error instanceof Error
            ? error.message
            : "Failed to check authentication status",
      });
    }
  };

  const signIn = async (email: string, password: string) => {
    setState((prev) => ({ ...prev, isLoading: true, error: null }));

    try {
      console.log("Attempting to sign in with:", email);
      const response = await api.auth.signin(email, password);
      console.log("Sign in response 2:", response);

      if (response.status === 200 && response.data) {
        console.log(
          "Auth successful, updating state with user:",
          response.data.user
        );

        setState({
          user: response.data.user,
          isLoading: false,
          isAuthenticated: true,
          error: null,
        });

        // Force a re-check of auth status after a short delay
        setTimeout(() => {
          console.log("Re-checking auth status after login");
          checkAuthStatus();
        }, 500);

        return true;
      } else {
        console.log("Auth failed:", response.error);
        setState((prev) => ({
          ...prev,
          isLoading: false,
          error: response.error || "Failed to sign in",
        }));
        return false;
      }
    } catch (error) {
      console.error("Sign in error:", error);
      setState((prev) => ({
        ...prev,
        isLoading: false,
        error: error instanceof Error ? error.message : "Failed to sign in",
      }));
      return false;
    }
  };

  const signUp = async (email: string, password: string) => {
    setState((prev) => ({ ...prev, isLoading: true, error: null }));

    try {
      const response = await api.auth.signup(email, password);

      if (response.status === 201 && response.data) {
        setState({
          user: response.data.user,
          isLoading: false,
          isAuthenticated: true,
          error: null,
        });
        return true;
      } else {
        setState((prev) => ({
          ...prev,
          isLoading: false,
          error: response.error || "Failed to sign up",
        }));
        return false;
      }
    } catch (error) {
      setState((prev) => ({
        ...prev,
        isLoading: false,
        error: error instanceof Error ? error.message : "Failed to sign up",
      }));
      return false;
    }
  };

  const signOut = async () => {
    setState((prev) => ({ ...prev, isLoading: true, error: null }));

    try {
      await api.auth.signout();

      setState({
        user: null,
        isLoading: false,
        isAuthenticated: false,
        error: null,
      });
      return true;
    } catch (error) {
      setState((prev) => ({
        ...prev,
        isLoading: false,
        error: error instanceof Error ? error.message : "Failed to sign out",
      }));
      return false;
    }
  };

  const resetPassword = async (email: string) => {
    setState((prev) => ({ ...prev, isLoading: true, error: null }));

    try {
      const response = await api.auth.resetPassword(email);

      setState((prev) => ({ ...prev, isLoading: false }));
      return response.status === 200;
    } catch (error) {
      setState((prev) => ({
        ...prev,
        isLoading: false,
        error:
          error instanceof Error ? error.message : "Failed to reset password",
      }));
      return false;
    }
  };

  useEffect(() => {
    checkAuthStatus();
  }, []);

  return {
    ...state,
    signIn,
    signUp,
    signOut,
    resetPassword,
  };
}
