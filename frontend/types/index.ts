/**
 * Type definitions for the Lilo outfit finder application
 */

// User and Authentication Types
export interface User {
  id: string;
  email: string;
  name?: string;
  picture?: string;
  styleProfile?: StyleProfile;
  createdAt: Date;
  updatedAt: Date;
}

export interface StyleProfile {
  id: string;
  userId: string;
  preferredStyles: string[];
  weeklySchedule: WeeklySchedule;
  seasonalPreferences: Record<string, string[]>;
  colorPreferences: string[];
  updatedAt: Date;
}

export type DayType =
  | "casual"
  | "professional"
  | "formal"
  | "athletic"
  | "other";

export interface WeeklySchedule {
  monday: DayType;
  tuesday: DayType;
  wednesday: DayType;
  thursday: DayType;
  friday: DayType;
  saturday: DayType;
  sunday: DayType;
}

// Wardrobe Types
export interface ClothingItem {
  id: string;
  userId: string;
  name: string;
  category: string;
  subcategory: string;
  color: string;
  season: string[];
  brand?: string;
  size?: string;
  imageUrls: string[];
  isOwned: boolean; // true for owned, false for wishlist
  createdAt: Date;
  updatedAt: Date;
}

// Outfit Types
export interface Outfit {
  id: string;
  userId: string;
  name: string;
  description?: string;
  items: string[]; // IDs of clothing items
  occasion: string[];
  season: string[];
  imageUrl?: string;
  isRecommended: boolean;
  isFavorite: boolean;
  createdAt: Date;
  updatedAt: Date;
}

// Reflection Types
export interface Reflection {
  id: string;
  userId: string;
  outfitId: string;
  date: Date;
  confidence: number; // 1-5 scale
  comfort: number; // 1-5 scale
  wouldRewear: boolean;
  notes?: string;
  createdAt: Date;
}

// Recommendation Types
export interface Recommendation {
  id: string;
  userId: string;
  outfitId: string;
  date: Date;
  feedback?: "liked" | "disliked" | "neutral";
  reason?: string;
  stylingTips: string[];
  createdAt: Date;
}

// API Response Types
export interface ApiResponse<T> {
  data?: T;
  error?: string;
  message?: string;
  status: number;
}
