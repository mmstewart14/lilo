/**
 * User model for the Lilo outfit finder application
 */
import { StyleProfile, User } from "../types";

export class UserModel {
  private user: User | null = null;
  private styleProfile: StyleProfile | null = null;

  constructor(userData?: User, styleProfileData?: StyleProfile) {
    if (userData) {
      this.user = userData;
    }
    if (styleProfileData) {
      this.styleProfile = styleProfileData;
    }
  }

  /**
   * Get user data
   */
  getUserData(): User | null {
    return this.user;
  }

  /**
   * Set user data
   */
  setUserData(userData: User): void {
    this.user = userData;
  }

  /**
   * Get style profile
   */
  getStyleProfile(): StyleProfile | null {
    return this.styleProfile;
  }

  /**
   * Set style profile
   */
  setStyleProfile(styleProfileData: StyleProfile): void {
    this.styleProfile = styleProfileData;
  }

  /**
   * Check if user has completed onboarding
   */
  hasCompletedOnboarding(): boolean {
    return !!this.styleProfile;
  }

  /**
   * Get user's preferred styles
   */
  getPreferredStyles(): string[] {
    return this.styleProfile?.preferredStyles || [];
  }

  /**
   * Get user's weekly schedule
   */
  getWeeklySchedule(): Record<string, string> | null {
    if (!this.styleProfile?.weeklySchedule) return null;

    // Convert WeeklySchedule to Record<string, string>
    const schedule = this.styleProfile.weeklySchedule;
    return {
      monday: schedule.monday,
      tuesday: schedule.tuesday,
      wednesday: schedule.wednesday,
      thursday: schedule.thursday,
      friday: schedule.friday,
      saturday: schedule.saturday,
      sunday: schedule.sunday,
    };
  }

  /**
   * Get day type for a specific day
   */
  getDayType(day: string): string | null {
    if (!this.styleProfile?.weeklySchedule) return null;

    const dayLower = day.toLowerCase();
    const schedule = this.styleProfile.weeklySchedule;

    if (dayLower === "monday") return schedule.monday;
    if (dayLower === "tuesday") return schedule.tuesday;
    if (dayLower === "wednesday") return schedule.wednesday;
    if (dayLower === "thursday") return schedule.thursday;
    if (dayLower === "friday") return schedule.friday;
    if (dayLower === "saturday") return schedule.saturday;
    if (dayLower === "sunday") return schedule.sunday;

    return null;
  }
}
