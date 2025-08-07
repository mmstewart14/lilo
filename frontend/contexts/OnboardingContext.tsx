/**
 * OnboardingContext for managing onboarding state across screens
 */
import AsyncStorage from "@react-native-async-storage/async-storage";
import {
  createContext,
  PropsWithChildren,
  useContext,
  useEffect,
  useState,
} from "react";
import api from "../services/api";
import { DayType, WeeklySchedule } from "../types";

// Define the style categories and options
export const STYLE_CATEGORIES = [
  "Casual",
  "Professional",
  "Formal",
  "Athletic",
  "Bohemian",
  "Vintage",
  "Minimalist",
  "Streetwear",
  "Preppy",
  "Edgy",
];

// Define the day types
export const DAY_TYPES: DayType[] = [
  "casual",
  "professional",
  "formal",
  "athletic",
  "other",
];

// Define the seasons
export const SEASONS = ["Spring", "Summer", "Fall", "Winter"];

// Define the colors
export const COLORS = [
  "Black",
  "White",
  "Gray",
  "Blue",
  "Navy",
  "Green",
  "Red",
  "Pink",
  "Purple",
  "Yellow",
  "Orange",
  "Brown",
  "Beige",
];

// Storage keys
const STORAGE_KEYS = {
  ONBOARDING_STATE: "lilo_onboarding_state",
  ONBOARDING_COMPLETED: "lilo_onboarding_completed",
  ONBOARDING_PROGRESS: "lilo_onboarding_progress",
};

// Define the onboarding state
interface OnboardingState {
  preferredStyles: string[];
  weeklySchedule: WeeklySchedule;
  seasonalPreferences: Record<string, string[]>;
  colorPreferences: string[];
  quickAddItems: {
    name: string;
    category: string;
    color: string;
    isOwned: boolean;
  }[];
  currentStep: number;
  completedSteps: number[];
  isOnboardingComplete: boolean;
}

// Define the context value
interface OnboardingContextValue {
  state: OnboardingState;
  setPreferredStyles: (styles: string[]) => void;
  setWeeklySchedule: (day: keyof WeeklySchedule, type: DayType) => void;
  setSeasonalPreferences: (season: string, styles: string[]) => void;
  setColorPreferences: (colors: string[]) => void;
  addQuickItem: (item: {
    name: string;
    category: string;
    color: string;
    isOwned: boolean;
  }) => void;
  removeQuickItem: (index: number) => void;
  nextStep: () => void;
  prevStep: () => void;
  goToStep: (step: number) => void;
  markStepCompleted: (step: number) => void;
  isStepCompleted: (step: number) => boolean;
  submitOnboarding: () => Promise<boolean>;
  resetOnboarding: () => void;
  saveOnboardingProgress: () => Promise<void>;
  loadOnboardingProgress: () => Promise<boolean>;
  completeOnboarding: () => Promise<void>;
  isOnboardingComplete: () => Promise<boolean>;
}

// Create the context
const OnboardingContext = createContext<OnboardingContextValue | undefined>(
  undefined
);

// Default weekly schedule
const defaultWeeklySchedule: WeeklySchedule = {
  monday: "professional",
  tuesday: "professional",
  wednesday: "professional",
  thursday: "professional",
  friday: "professional",
  saturday: "casual",
  sunday: "casual",
};

// Initial state
const initialState: OnboardingState = {
  preferredStyles: [],
  weeklySchedule: { ...defaultWeeklySchedule },
  seasonalPreferences: {
    Spring: [],
    Summer: [],
    Fall: [],
    Winter: [],
  },
  colorPreferences: [],
  quickAddItems: [],
  currentStep: 1,
  completedSteps: [],
  isOnboardingComplete: false,
};

// Provider component
export function OnboardingProvider({ children }: PropsWithChildren) {
  const [state, setState] = useState<OnboardingState>(initialState);
  const [isLoading, setIsLoading] = useState(true);

  // Load saved onboarding state on mount
  useEffect(() => {
    const loadSavedState = async () => {
      try {
        await loadOnboardingProgress();
      } catch (error) {
        console.error("Failed to load onboarding state:", error);
      } finally {
        setIsLoading(false);
      }
    };

    loadSavedState();
  }, []);

  const setPreferredStyles = (styles: string[]) => {
    setState((prev) => ({
      ...prev,
      preferredStyles: styles,
    }));
  };

  const setWeeklySchedule = (day: keyof WeeklySchedule, type: DayType) => {
    setState((prev) => ({
      ...prev,
      weeklySchedule: {
        ...prev.weeklySchedule,
        [day]: type,
      },
    }));
  };

  const setSeasonalPreferences = (season: string, styles: string[]) => {
    setState((prev) => ({
      ...prev,
      seasonalPreferences: {
        ...prev.seasonalPreferences,
        [season]: styles,
      },
    }));
  };

  const setColorPreferences = (colors: string[]) => {
    setState((prev) => ({
      ...prev,
      colorPreferences: colors,
    }));
  };

  const addQuickItem = (item: {
    name: string;
    category: string;
    color: string;
    isOwned: boolean;
  }) => {
    setState((prev) => ({
      ...prev,
      quickAddItems: [...prev.quickAddItems, item],
    }));
  };

  const removeQuickItem = (index: number) => {
    setState((prev) => ({
      ...prev,
      quickAddItems: prev.quickAddItems.filter((_, i) => i !== index),
    }));
  };

  const nextStep = () => {
    setState((prev) => {
      const nextStepNumber = prev.currentStep + 1;
      // Mark the current step as completed if it's not already
      const updatedCompletedSteps = prev.completedSteps.includes(
        prev.currentStep
      )
        ? prev.completedSteps
        : [...prev.completedSteps, prev.currentStep];

      return {
        ...prev,
        currentStep: nextStepNumber,
        completedSteps: updatedCompletedSteps,
      };
    });

    // Save progress after step change
    saveOnboardingProgress();
  };

  const prevStep = () => {
    setState((prev) => ({
      ...prev,
      currentStep: Math.max(1, prev.currentStep - 1),
    }));
  };

  const goToStep = (step: number) => {
    if (step < 1) return;
    setState((prev) => ({
      ...prev,
      currentStep: step,
    }));
  };

  const markStepCompleted = (step: number) => {
    setState((prev) => {
      if (prev.completedSteps.includes(step)) {
        return prev;
      }
      return {
        ...prev,
        completedSteps: [...prev.completedSteps, step],
      };
    });

    // Save progress after marking step as completed
    saveOnboardingProgress();
  };

  const isStepCompleted = (step: number): boolean => {
    return state.completedSteps.includes(step);
  };

  const saveOnboardingProgress = async (): Promise<void> => {
    try {
      const stateToSave = JSON.stringify(state);
      await AsyncStorage.setItem(STORAGE_KEYS.ONBOARDING_STATE, stateToSave);
      await AsyncStorage.setItem(
        STORAGE_KEYS.ONBOARDING_PROGRESS,
        JSON.stringify(state.completedSteps)
      );
    } catch (error) {
      console.error("Failed to save onboarding progress:", error);
    }
  };

  const loadOnboardingProgress = async (): Promise<boolean> => {
    try {
      const savedState = await AsyncStorage.getItem(
        STORAGE_KEYS.ONBOARDING_STATE
      );
      const isCompleted = await AsyncStorage.getItem(
        STORAGE_KEYS.ONBOARDING_COMPLETED
      );

      if (savedState) {
        const parsedState = JSON.parse(savedState);
        setState({
          ...parsedState,
          isOnboardingComplete: isCompleted === "true",
        });
        return true;
      }
      return false;
    } catch (error) {
      console.error("Failed to load onboarding progress:", error);
      return false;
    }
  };

  const completeOnboarding = async (): Promise<void> => {
    try {
      await AsyncStorage.setItem(STORAGE_KEYS.ONBOARDING_COMPLETED, "true");
      setState((prev) => ({
        ...prev,
        isOnboardingComplete: true,
      }));
    } catch (error) {
      console.error("Failed to mark onboarding as complete:", error);
    }
  };

  const isOnboardingComplete = async (): Promise<boolean> => {
    try {
      const isCompleted = await AsyncStorage.getItem(
        STORAGE_KEYS.ONBOARDING_COMPLETED
      );
      return isCompleted === "true";
    } catch (error) {
      console.error("Failed to check if onboarding is complete:", error);
      return false;
    }
  };

  const submitOnboarding = async (): Promise<boolean> => {
    try {
      // Create style profile
      const styleProfileData = {
        preferredStyles: state.preferredStyles,
        weeklySchedule: state.weeklySchedule,
        seasonalPreferences: state.seasonalPreferences,
        colorPreferences: state.colorPreferences,
      };

      const response = await api.users.updateStyleProfile(styleProfileData);

      if (response.error) {
        console.error("Failed to update style profile:", response.error);
        return false;
      }

      // Add quick items if any
      if (state.quickAddItems.length > 0) {
        for (const item of state.quickAddItems) {
          const itemData = {
            name: item.name,
            category: item.category,
            subcategory: "",
            color: item.color,
            season: ["Spring", "Summer", "Fall", "Winter"], // Default to all seasons
            isOwned: item.isOwned,
            imageUrls: [],
          };

          const itemResponse = await api.wardrobe.addItem(itemData);
          if (itemResponse.error) {
            console.error("Failed to add item:", itemResponse.error);
            // Continue with other items even if one fails
          }
        }
      }

      // Mark onboarding as complete
      await completeOnboarding();

      return true;
    } catch (error) {
      console.error("Error submitting onboarding data:", error);
      return false;
    }
  };

  const resetOnboarding = () => {
    // Clear storage
    AsyncStorage.multiRemove([
      STORAGE_KEYS.ONBOARDING_STATE,
      STORAGE_KEYS.ONBOARDING_COMPLETED,
      STORAGE_KEYS.ONBOARDING_PROGRESS,
    ]).catch((error) => {
      console.error("Failed to clear onboarding storage:", error);
    });

    // Reset state
    setState({
      ...initialState,
      currentStep: 1,
      completedSteps: [],
      isOnboardingComplete: false,
    });
  };

  // Skip rendering until initial loading is complete
  if (isLoading) {
    return null;
  }

  const value: OnboardingContextValue = {
    state,
    setPreferredStyles,
    setWeeklySchedule,
    setSeasonalPreferences,
    setColorPreferences,
    addQuickItem,
    removeQuickItem,
    nextStep,
    prevStep,
    goToStep,
    markStepCompleted,
    isStepCompleted,
    submitOnboarding,
    resetOnboarding,
    saveOnboardingProgress,
    loadOnboardingProgress,
    completeOnboarding,
    isOnboardingComplete,
  };

  return (
    <OnboardingContext.Provider value={value}>
      {children}
    </OnboardingContext.Provider>
  );
}

// Custom hook to use the onboarding context
export function useOnboarding() {
  const context = useContext(OnboardingContext);
  if (context === undefined) {
    throw new Error("useOnboarding must be used within an OnboardingProvider");
  }
  return context;
}
