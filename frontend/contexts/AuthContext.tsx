/**
 * Authentication Context Provider for the Lilo outfit finder application
 * Provides global authentication state management
 */
import {
  createContext,
  ReactNode,
  useContext,
  useEffect,
  useState,
} from "react";
import api from "../services/api";
import { User } from "../types";

interface AuthState {
  user: User | null;
  isLoading: boolean;
  isAuthenticated: boolean;
  error: string | null;
}

interface AuthContextValue extends AuthState {
  signIn: (email: string, password: string) => Promise<boolean>;
  signUp: (email: string, password: string) => Promise<boolean>;
  signOut: () => Promise<boolean>;
  resetPassword: (email: string) => Promise<boolean>;
  refreshUser: () => Promise<void>;

  // TODO: Initialize session maybe? https://github.com/orgs/supabase/discussions/20155
}

const AuthContext = createContext<AuthContextValue | undefined>(undefined);

export function AuthProvider({ children }: { children: ReactNode }) {
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

  const signIn = async (email: string, password: string): Promise<boolean> => {
    setState((prev) => ({ ...prev, isLoading: true, error: null }));

    try {
      console.log("Attempting to sign in with:", email);
      const response = await api.auth.signin(email, password);
      console.log("Sign in response:", response);

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

        // Force a re-check of auth status after a short delay to get complete user data
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

  const signUp = async (email: string, password: string): Promise<boolean> => {
    setState((prev) => ({ ...prev, isLoading: true, error: null }));

    try {
      console.log("Attempting to sign up with:", email);
      const response = await api.auth.signup(email, password);
      console.log("Sign up response:", response);

      if (response.status === 201 && response.data) {
        setState({
          user: response.data.user,
          isLoading: false,
          isAuthenticated: true,
          error: null,
        });
        return true;
      } else {
        console.log("Sign up failed:", response.error);
        setState((prev) => ({
          ...prev,
          isLoading: false,
          error: response.error || "Failed to sign up",
        }));
        return false;
      }
    } catch (error) {
      console.log("Sign up error:", error);
      setState((prev) => ({
        ...prev,
        isLoading: false,
        error: error instanceof Error ? error.message : "Failed to sign up",
      }));
      return false;
    }
  };

  const signOut = async (): Promise<boolean> => {
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

  const resetPassword = async (email: string): Promise<boolean> => {
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

  const refreshUser = async (): Promise<void> => {
    await checkAuthStatus();
  };

  useEffect(() => {
    checkAuthStatus();
  }, []);

  const value: AuthContextValue = {
    ...state,
    signIn,
    signUp,
    signOut,
    resetPassword,
    refreshUser,
  };

  return <AuthContext.Provider value={value}>{children}</AuthContext.Provider>;
}

export function useAuth(): AuthContextValue {
  const context = useContext(AuthContext);
  if (context === undefined) {
    throw new Error("useAuth must be used within an AuthProvider");
  }
  return context;
}
