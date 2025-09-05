/**
 * Backend API client for the Lilo application
 * Handles all communication with the Go backend
 */
import { supabaseClient } from "./supabaseClient";
import { devLog, devError } from "../utils/devLog";
import { ApiResponse, StyleProfile, User } from "@/types";

// Dynamic base URL based on environment
const getBaseURL = (): string => {
  const envURL = process.env.EXPO_PUBLIC_API_BASE_URL;

  if (envURL) {
    return envURL;
  }

  // Fallback logic
  if (__DEV__) {
    // In development, try to detect if we're on a physical device
    // This is a simple heuristic - you might want to make this more sophisticated
    return "http://192.168.1.63:8080"; // Your computer's IP
  }

  return "http://localhost:8080"; // Default for web/simulator
};

const BASE_URL = getBaseURL();

/**
 * Generic request function for backend API calls
 */
async function request<T>(
  endpoint: string,
  method: "GET" | "POST" | "PUT" | "DELETE" = "GET",
  data?: any,
  requiresAuth: boolean = true
): Promise<ApiResponse<T>> {
  try {
    const url = `${BASE_URL}${endpoint}`;

    const headers: Record<string, string> = {
      "Content-Type": "application/json",
    };

    // Add auth header if required
    if (requiresAuth) {
      const { data: sessionData, error: sessionError } =
        await supabaseClient.auth.getSession();
      const token = sessionData?.session?.access_token;

      devLog("Auth session check:", {
        hasSession: !!sessionData?.session,
        hasToken: !!token,
        tokenPreview: token ? `${token.substring(0, 20)}...` : null,
        sessionError,
        expiresAt: sessionData?.session?.expires_at,
        user: sessionData?.session?.user?.email,
      });

      if (!token) {
        devError("No auth token available for backend request");
        return {
          error: "Authentication required",
          status: 401,
        };
      }

      headers.Authorization = `Bearer ${token}`;
    }

    const options: RequestInit = {
      method,
      headers,
    };

    if (data && (method === "POST" || method === "PUT")) {
      options.body = JSON.stringify(data);
    }

    devLog(`SENDING REQUEST -- ${method} ${url}`, data ?? "");

    const response = await fetch(url, options);
    const responseData = await response.json();

    devLog(`RECEIVED RESPONSE -- ${method} ${url}`, {
      status: response.status,
      data: responseData,
    });

    return {
      data: responseData.data,
      message: responseData.message,
      error: responseData.error,
      status: response.status,
    };
  } catch (error) {
    devError("Backend request failed:", error);
    return {
      error: error instanceof Error ? error.message : "Unknown error occurred",
      status: 500,
    };
  }
}

// Health check
export const healthCheck = (): Promise<ApiResponse<any>> =>
  request("/api/health", "GET", undefined, false);

// User endpoints
export const user = {
  createProfile: (data: { supabaseId: string; email: string }) =>
    request<User>("/api/auth/signup", "POST", data),

  getProfile: () => request<User>("/api/auth/user", "GET"),

  updateProfile: (data: any) =>
    request<User>("/api/users/profile", "PUT", data),

  getStyleProfile: () =>
    request<StyleProfile>("/api/users/style-profile", "GET"),

  updateStyleProfile: (data: any) =>
    request<StyleProfile>("/api/users/style-profile", "PUT", data),
};

// Wardrobe endpoints
export const wardrobe = {
  getItems: (filters?: any) => request("/api/wardrobe/items", "GET", filters),

  addItem: (item: any) => request("/api/wardrobe/items", "POST", item),

  getItem: (id: string) => request(`/api/wardrobe/items/${id}`, "GET"),

  updateItem: (id: string, data: any) =>
    request(`/api/wardrobe/items/${id}`, "PUT", data),

  deleteItem: (id: string) => request(`/api/wardrobe/items/${id}`, "DELETE"),

  getCategories: () => request("/api/wardrobe/categories", "GET"),
};

// Outfit endpoints
export const outfits = {
  getOutfits: (filters?: any) => request("/api/outfits", "GET", filters),

  createOutfit: (outfit: any) => request("/api/outfits", "POST", outfit),

  getOutfit: (id: string) => request(`/api/outfits/${id}`, "GET"),

  updateOutfit: (id: string, data: any) =>
    request(`/api/outfits/${id}`, "PUT", data),

  deleteOutfit: (id: string) => request(`/api/outfits/${id}`, "DELETE"),

  favoriteOutfit: (id: string) =>
    request(`/api/outfits/${id}/favorite`, "POST"),

  unfavoriteOutfit: (id: string) =>
    request(`/api/outfits/${id}/favorite`, "DELETE"),
};

// Recommendation endpoints
export const recommendations = {
  getDaily: () => request("/api/recommendations/daily", "GET"),

  getExplore: (filters?: any) =>
    request("/api/recommendations/explore", "GET", filters),

  submitFeedback: (recommendationId: string, feedback: string) =>
    request("/api/recommendations/feedback", "POST", {
      recommendationId,
      feedback,
    }),
};

// Reflection endpoints
export const reflections = {
  submitReflection: (data: any) => request("/api/reflections", "POST", data),

  getReflections: () => request("/api/reflections", "GET"),

  getInsights: () => request("/api/reflections/insights", "GET"),
};

// Wishlist endpoints
export const wishlist = {
  getItems: () => request("/api/wishlist", "GET"),

  addItem: (itemId: string) => request("/api/wishlist", "POST", { itemId }),

  removeItem: (id: string) => request(`/api/wishlist/${id}`, "DELETE"),
};

// Export all endpoints as a single object for convenience
export const backendClient = {
  healthCheck,
  user,
  wardrobe,
  outfits,
  recommendations,
  reflections,
  wishlist,
};

export default backendClient;
