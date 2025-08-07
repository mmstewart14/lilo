/**
 * API service for the Lilo outfit finder application
 */
import { ApiResponse, User } from "../types";
import { supabase } from "./supabase";

/**
 * Generic API request function
 */
async function request<T>(
  endpoint: string,
  method: "GET" | "POST" | "PUT" | "DELETE" = "GET",
  data?: any,
  headers?: Record<string, string>
): Promise<ApiResponse<T>> {
  try {
    console.log(endpoint);

    // Get the current session for authentication
    const { data: sessionData } = await supabase.auth.getSession();
    const token = sessionData?.session?.access_token;

    console.log(sessionData);

    // Add authorization header if token exists
    const authHeaders = token
      ? { Authorization: `Bearer ${token}`, ...headers }
      : headers;

    const url = `${process.env.EXPO_PUBLIC_API_BASE_URL}${endpoint}`;
    const options: RequestInit = {
      method,
      headers: {
        "Content-Type": "application/json",
        ...authHeaders,
      },
    };

    if (data) {
      options.body = JSON.stringify(data);
    }

    const response = await fetch(url, options);

    const responseData = await response.json();

    console.log(responseData);

    return {
      data: responseData.data,
      message: responseData.message,
      error: responseData.error,
      status: response.status,
    };
  } catch (error) {
    console.error(
      "API request failed:",
      error instanceof Error ? error.message : error
    );
    return {
      error: error instanceof Error ? error.message : "Unknown error occurred",
      status: 500,
    };
  }
}

/**
 * API service object with methods for different endpoints
 */
const api = {
  // Auth endpoints using Supabase
  auth: {
    signup: async (email: string, password: string) => {
      const { data, error } = await supabase.auth.signUp({
        email,
        password,
      });

      if (error) {
        return {
          error: error.message,
          status: 400,
        };
      }

      // Create user profile in our backend
      if (data.user) {
        await request("/api/auth/signup", "POST", {
          supabaseId: data.user.id,
          email: data.user.email,
        });
      }

      return {
        data: {
          token: data.session?.access_token,
          user: mapSupabaseUser(data.user),
        },
        status: 201,
      };
    },

    signin: async (email: string, password: string) => {
      const { data, error } = await supabase.auth.signInWithPassword({
        email,
        password,
      });

      if (error) {
        return {
          error: error.message,
          status: 401,
        };
      }

      return {
        data: {
          token: data.session?.access_token,
          user: mapSupabaseUser(data.user),
        },
        status: 200,
      };
    },

    signout: async () => {
      const { error } = await supabase.auth.signOut();

      if (error) {
        return {
          error: error.message,
          status: 500,
        };
      }

      return {
        status: 200,
      };
    },

    getUser: async () => {
      const { data, error } = await supabase.auth.getUser();

      if (error || !data.user) {
        return {
          error: error?.message || "User not found",
          status: 401,
        };
      }

      // Get additional user data from our backend
      const userResponse = await request<User>("/api/users/profile", "GET");

      if (userResponse.error) {
        return {
          data: mapSupabaseUser(data.user),
          status: 200,
        };
      }

      return {
        data: {
          ...mapSupabaseUser(data.user),
          ...userResponse.data,
        },
        status: 200,
      };
    },

    resetPassword: async (email: string) => {
      const { error } = await supabase.auth.resetPasswordForEmail(email);

      if (error) {
        return {
          error: error.message,
          status: 400,
        };
      }

      return {
        message: "Password reset email sent",
        status: 200,
      };
    },
  },

  // User profile endpoints
  users: {
    getProfile: () => request("/api/users/profile", "GET"),
    updateProfile: (data: any) => request("/api/users/profile", "PUT", data),
    getStyleProfile: () => request("/api/users/style-profile", "GET"),
    updateStyleProfile: (data: any) =>
      request("/api/users/style-profile", "PUT", data),
  },

  // Wardrobe endpoints
  wardrobe: {
    getItems: (filters?: any) => request("/api/wardrobe/items", "GET", filters),
    addItem: (item: any) => request("/api/wardrobe/items", "POST", item),
    getItem: (id: string) => request(`/api/wardrobe/items/${id}`, "GET"),
    updateItem: (id: string, data: any) =>
      request(`/api/wardrobe/items/${id}`, "PUT", data),
    deleteItem: (id: string) => request(`/api/wardrobe/items/${id}`, "DELETE"),
    getCategories: () => request("/api/wardrobe/categories", "GET"),
  },

  // Outfit endpoints
  outfits: {
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
  },

  // Recommendation endpoints
  recommendations: {
    getDaily: () => request("/api/recommendations/daily", "GET"),
    getExplore: (filters?: any) =>
      request("/api/recommendations/explore", "GET", filters),
    submitFeedback: (recommendationId: string, feedback: string) =>
      request("/api/recommendations/feedback", "POST", {
        recommendationId,
        feedback,
      }),
  },

  // Reflection endpoints
  reflections: {
    submitReflection: (data: any) => request("/api/reflections", "POST", data),
    getReflections: () => request("/api/reflections", "GET"),
    getInsights: () => request("/api/reflections/insights", "GET"),
  },

  // Wishlist endpoints
  wishlist: {
    getItems: () => request("/api/wishlist", "GET"),
    addItem: (itemId: string) => request("/api/wishlist", "POST", { itemId }),
    removeItem: (id: string) => request(`/api/wishlist/${id}`, "DELETE"),
  },
};

/**
 * Helper function to map Supabase user to our User type
 */
function mapSupabaseUser(supabaseUser: any): User {
  if (!supabaseUser) return null as any;

  return {
    id: supabaseUser.id,
    email: supabaseUser.email || "",
    name: supabaseUser.user_metadata?.name || "",
    picture: supabaseUser.user_metadata?.avatar_url || "",
    createdAt: new Date(supabaseUser.created_at),
    updatedAt: new Date(supabaseUser.updated_at || supabaseUser.created_at),
  };
}

export default api;
