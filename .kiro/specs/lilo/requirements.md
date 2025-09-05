# Requirements Document

## Introduction

Lilo is a mobile application built with React Native and a Go backend that delivers personalized daily outfit inspiration based on a user's unique wardrobe, style preferences, and schedule. It combines reactive personalization with user-driven exploration, enabling wardrobe reflection, wishlist building, and outfit journaling — all wrapped in an intuitive and emotionally engaging experience. The system aims to simplify the daily decision-making process of choosing what to wear while helping users discover new outfit combinations.

## Requirements

### 1. User Authentication & Onboarding

**User Story:** As a new user, I want to securely authenticate and complete a personalized onboarding process, so that I can receive outfit recommendations tailored to my style and schedule.

#### Acceptance Criteria

1. WHEN a new user opens the app THEN the system SHALL provide Supabase Auth authentication options
2. WHEN a user authenticates with Supabase Auth THEN the system SHALL securely handle the authentication flow
3. WHEN a user completes authentication THEN the system SHALL present a short onboarding survey
4. WHEN completing the survey THEN the system SHALL collect information about desired style(s)
5. WHEN completing the survey THEN the system SHALL collect information about weekly outfit schedule (casual vs. professional days)
6. WHEN completing the survey THEN the system SHALL allow optional clothing inventory input
7. WHEN the survey is completed THEN the system SHALL generate a personalized style profile
8. IF a user wants to skip onboarding elements THEN the system SHALL allow them to do so
9. WHEN a user wants to revisit onboarding THEN the system SHALL provide access to previously completed elements

### 2. Daily Outfit Inspiration

**User Story:** As a user, I want to receive daily outfit inspiration, so that I can discover new outfit combinations and simplify my clothing decisions.

#### Acceptance Criteria

1. WHEN a user opens the app daily THEN the system SHALL present a unique outfit inspiration image
2. WHEN viewing the daily inspiration THEN the system SHALL display a breakdown of worn items
3. WHEN viewing the daily inspiration THEN the system SHALL provide styling tips
4. WHEN presented with an outfit THEN the system SHALL allow users to indicate whether they'd wear it or not
5. WHEN a user provides feedback on an outfit THEN the system SHALL use this data to influence future suggestions
6. IF a user doesn't interact with a daily outfit THEN the system SHALL remind them via notification

### 3. Owned vs. Aspirational Mode

**User Story:** As a user, I want to toggle between seeing outfits with only items I own and inspirational outfits with items I might not own, so that I can both utilize my current wardrobe and get ideas for future purchases.

#### Acceptance Criteria

1. WHEN browsing outfits THEN the system SHALL provide options to view outfits using only owned items
2. WHEN browsing outfits THEN the system SHALL provide options to view inspirational outfits (including unowned items)
3. WHEN displaying outfits with unowned items THEN the system SHALL clearly mark these items
4. WHEN viewing unowned items THEN the system SHALL allow users to add them to a personal Wishlist
5. WHEN a user adds an item to their Wishlist THEN the system SHALL save this preference for future reference

### 4. Explore Page (Free Browsing)

**User Story:** As a user, I want to freely browse a gallery of outfits with various filtering options, so that I can find inspiration based on specific criteria.

#### Acceptance Criteria

1. WHEN a user accesses the Explore page THEN the system SHALL display a scrollable gallery of outfits
2. WHEN browsing outfits THEN the system SHALL provide filters for style (e.g., trendy, classic)
3. WHEN browsing outfits THEN the system SHALL provide filters for season (e.g., fall, summer)
4. WHEN browsing outfits THEN the system SHALL provide filters for formality (e.g., work, lounge)
5. WHEN browsing outfits THEN the system SHALL provide filters for specific clothing items (e.g., "black boots")
6. WHEN a user applies filters THEN the system SHALL update results in real-time
7. IF search functionality is implemented THEN the system SHALL allow users to search by keyword or item

### 5. End-of-Day Outfit Reflection

**User Story:** As a user, I want to reflect on my daily outfit choices, so that I can improve future outfit selections and build a collection of favorites.

#### Acceptance Criteria

1. WHEN the day ends THEN the system SHALL send a push notification asking for outfit reflection
2. WHEN prompted for reflection THEN the system SHALL ask "How did you feel in today's outfit?"
3. WHEN reflecting on an outfit THEN the system SHALL allow users to rate confidence/comfort
4. WHEN reflecting on an outfit THEN the system SHALL allow users to indicate rewear likelihood
5. WHEN reflecting on an outfit THEN the system SHALL provide an option to save the outfit
6. WHEN a user provides reflection data THEN the system SHALL use this to improve personalization
7. WHEN a user positively rates an outfit THEN the system SHALL offer to save it as a favorite

### 6. Technical & UX Considerations

**User Story:** As a user, I want an intuitive, beautiful, and responsive app experience, so that using the app feels enjoyable and effortless.

#### Acceptance Criteria

1. WHEN using the app THEN the system SHALL provide a simple, beautiful UX with strong emotional tone
2. WHEN interacting with the app THEN the system SHALL respond with fast interactions
3. WHEN browsing outfits THEN the system SHALL emphasize visual browsing
4. WHEN the app needs to store user data THEN the system SHALL save style profile, clothing inventory, outfit feedback history, wishlist, and saved outfits
5. WHEN generating recommendations THEN the system SHALL adapt over time based on explicit (thumbs up/down) and reflective feedback
6. WHEN engagement is needed THEN the system SHALL use push notifications for habit-building
7. IF the app is used on different platforms THEN the system SHALL provide consistent experience on both iOS and Android

### 7. Go Backend API

**User Story:** As a developer, I want a robust Go backend API, so that the mobile application can securely store and retrieve data.

#### Acceptance Criteria

1. WHEN the mobile app makes API requests THEN the system SHALL authenticate and authorize the requests
2. WHEN storing user data THEN the system SHALL encrypt sensitive information
3. WHEN handling concurrent requests THEN the system SHALL maintain data consistency
4. WHEN processing image uploads THEN the system SHALL optimize and store images efficiently
5. WHEN generating outfit recommendations THEN the system SHALL use efficient algorithms
6. IF the API experiences errors THEN the system SHALL log details and return appropriate error responses
7. WHEN the system needs to scale THEN the Go backend SHALL handle increased load efficiently

### 8. Performance and Reliability

**User Story:** As a user, I want the app to be fast and reliable, so that I can use it without frustration.

#### Acceptance Criteria

1. WHEN a user opens the app THEN the system SHALL load within 3 seconds on standard connections
2. WHEN a user navigates between screens THEN the system SHALL respond within 1 second
3. WHEN a user is offline THEN the system SHALL provide access to previously loaded data
4. WHEN connection is restored THEN the system SHALL synchronize local changes with the server
5. WHEN the app is in the background THEN the system SHALL minimize battery consumption
6. IF the app crashes THEN the system SHALL collect diagnostic information for developers
7. WHEN the app updates THEN the system SHALL preserve user data and settings
