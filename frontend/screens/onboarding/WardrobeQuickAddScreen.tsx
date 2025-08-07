/**
 * WardrobeQuickAddScreen for the onboarding flow
 * Allows users to quickly add some initial clothing items to their wardrobe
 */
import { useState } from "react";
import {
  Alert,
  FlatList,
  StyleSheet,
  TextInput,
  TouchableOpacity,
  View,
} from "react-native";
import { ThemedText } from "../../components/ThemedText";
import { ThemedView } from "../../components/ThemedView";
import { ClothingItemCard } from "../../components/wardrobe/ClothingItemCard";
import { COLORS, useOnboarding } from "../../contexts/OnboardingContext";
import { useThemeColor } from "../../hooks/useThemeColor";
import { OnboardingLayout } from "./OnboardingLayout";

interface WardrobeQuickAddScreenProps {
  onNext: () => void;
  onBack: () => void;
  onSkip: () => void;
}

// Define clothing categories
const CLOTHING_CATEGORIES = [
  "Tops",
  "Bottoms",
  "Dresses",
  "Outerwear",
  "Shoes",
  "Accessories",
];

export function WardrobeQuickAddScreen({
  onNext,
  onBack,
  onSkip,
}: WardrobeQuickAddScreenProps) {
  const { state, addQuickItem, removeQuickItem } = useOnboarding();
  const { quickAddItems } = state;

  const [itemName, setItemName] = useState("");
  const [selectedCategory, setSelectedCategory] = useState(
    CLOTHING_CATEGORIES[0]
  );
  const [selectedColor, setSelectedColor] = useState(COLORS[0]);
  const [isWishlist, setIsWishlist] = useState(false);

  const primaryColor = useThemeColor(
    { light: "#6200ee", dark: "#bb86fc" },
    "primary"
  );
  const cardBgColor = useThemeColor(
    { light: "#ffffff", dark: "#1e1e1e" },
    "card"
  );
  const borderColor = useThemeColor(
    { light: "#e0e0e0", dark: "#333333" },
    "border"
  );
  const textInputBgColor = useThemeColor(
    { light: "#f5f5f5", dark: "#2a2a2a" },
    "inputBackground"
  );
  const textColor = useThemeColor(
    { light: "#000000", dark: "#ffffff" },
    "text"
  );

  const handleAddItem = () => {
    if (!itemName.trim()) {
      Alert.alert("Error", "Please enter an item name");
      return;
    }

    addQuickItem({
      name: itemName.trim(),
      category: selectedCategory,
      color: selectedColor,
      isOwned: !isWishlist,
    });

    // Reset form
    setItemName("");
    setIsWishlist(false);
  };

  const handleDeleteItem = (index: number) => {
    Alert.alert("Delete Item", "Are you sure you want to remove this item?", [
      {
        text: "Cancel",
        style: "cancel",
      },
      {
        text: "Delete",
        onPress: () => removeQuickItem(index),
        style: "destructive",
      },
    ]);
  };

  const handleNext = () => {
    onNext();
  };

  const renderQuickAddItem = ({
    item,
    index,
  }: {
    item: any;
    index: number;
  }) => {
    // Create a mock clothing item for the card
    const mockItem = {
      id: `temp-${index}`,
      userId: "",
      name: item.name,
      category: item.category,
      subcategory: "",
      color: item.color,
      season: ["Spring", "Summer", "Fall", "Winter"],
      imageUrls: [],
      isOwned: item.isOwned,
      createdAt: new Date(),
      updatedAt: new Date(),
    };

    return (
      <ClothingItemCard
        item={mockItem}
        onDelete={() => handleDeleteItem(index)}
      />
    );
  };

  return (
    <OnboardingLayout
      title="Quick Add Items"
      subtitle="Add a few key items from your wardrobe to get started. You can add more later."
      currentStep={3}
      totalSteps={3}
      onNext={handleNext}
      onBack={onBack}
      onSkip={onSkip}
      nextLabel="Finish"
      showSkip={false}
    >
      <ThemedView style={styles.formContainer}>
        <TextInput
          style={[
            styles.input,
            { backgroundColor: textInputBgColor, color: textColor },
          ]}
          placeholder="Item name (e.g. Blue Jeans)"
          placeholderTextColor="#999"
          value={itemName}
          onChangeText={setItemName}
        />

        <ThemedText style={styles.sectionTitle}>Category</ThemedText>
        <View style={styles.optionsContainer}>
          {CLOTHING_CATEGORIES.map((category) => (
            <TouchableOpacity
              key={category}
              style={[
                styles.optionButton,
                {
                  backgroundColor:
                    selectedCategory === category ? primaryColor : cardBgColor,
                  borderColor:
                    selectedCategory === category ? primaryColor : borderColor,
                },
              ]}
              onPress={() => setSelectedCategory(category)}
            >
              <ThemedText
                style={[
                  styles.optionText,
                  {
                    color: selectedCategory === category ? "white" : undefined,
                  },
                ]}
              >
                {category}
              </ThemedText>
            </TouchableOpacity>
          ))}
        </View>

        <ThemedText style={styles.sectionTitle}>Color</ThemedText>
        <View style={styles.optionsContainer}>
          {COLORS.map((color) => (
            <TouchableOpacity
              key={color}
              style={[
                styles.colorButton,
                {
                  backgroundColor:
                    selectedColor === color ? primaryColor : cardBgColor,
                  borderColor:
                    selectedColor === color ? primaryColor : borderColor,
                },
              ]}
              onPress={() => setSelectedColor(color)}
            >
              <ThemedText
                style={[
                  styles.optionText,
                  {
                    color: selectedColor === color ? "white" : undefined,
                  },
                ]}
              >
                {color}
              </ThemedText>
            </TouchableOpacity>
          ))}
        </View>

        <View style={styles.wishlistContainer}>
          <TouchableOpacity
            style={[
              styles.wishlistButton,
              {
                backgroundColor: isWishlist ? primaryColor : cardBgColor,
                borderColor: isWishlist ? primaryColor : borderColor,
              },
            ]}
            onPress={() => setIsWishlist(!isWishlist)}
          >
            <ThemedText
              style={[
                styles.wishlistText,
                {
                  color: isWishlist ? "white" : undefined,
                },
              ]}
            >
              Add to Wishlist
            </ThemedText>
          </TouchableOpacity>
          <ThemedText style={styles.wishlistHelp}>
            Toggle if you dont own this item yet
          </ThemedText>
        </View>

        <TouchableOpacity
          style={[styles.addButton, { backgroundColor: primaryColor }]}
          onPress={handleAddItem}
        >
          <ThemedText style={styles.addButtonText}>Add Item</ThemedText>
        </TouchableOpacity>
      </ThemedView>

      {quickAddItems.length > 0 && (
        <>
          <ThemedText style={styles.itemsTitle}>
            Added Items ({quickAddItems.length})
          </ThemedText>
          <FlatList
            data={quickAddItems}
            renderItem={renderQuickAddItem}
            keyExtractor={(_, index) => `item-${index}`}
            scrollEnabled={false}
          />
        </>
      )}

      <ThemedText style={styles.helperText}>
        You can add more items to your wardrobe later.
      </ThemedText>
    </OnboardingLayout>
  );
}

const styles = StyleSheet.create({
  formContainer: {
    marginBottom: 24,
  },
  input: {
    height: 50,
    borderRadius: 8,
    paddingHorizontal: 16,
    marginBottom: 16,
    fontSize: 16,
  },
  sectionTitle: {
    fontSize: 16,
    fontWeight: "600",
    marginBottom: 8,
  },
  optionsContainer: {
    flexDirection: "row",
    flexWrap: "wrap",
    marginBottom: 16,
    marginLeft: -4,
    marginRight: -4,
  },
  optionButton: {
    paddingHorizontal: 12,
    paddingVertical: 8,
    borderRadius: 20,
    marginHorizontal: 4,
    marginBottom: 8,
    borderWidth: 1,
  },
  colorButton: {
    paddingHorizontal: 12,
    paddingVertical: 8,
    borderRadius: 20,
    marginHorizontal: 4,
    marginBottom: 8,
    borderWidth: 1,
  },
  optionText: {
    fontSize: 14,
    fontWeight: "500",
  },
  wishlistContainer: {
    flexDirection: "row",
    alignItems: "center",
    marginBottom: 16,
  },
  wishlistButton: {
    paddingHorizontal: 12,
    paddingVertical: 8,
    borderRadius: 20,
    borderWidth: 1,
    marginRight: 8,
  },
  wishlistText: {
    fontSize: 14,
    fontWeight: "500",
  },
  wishlistHelp: {
    fontSize: 12,
    opacity: 0.6,
  },
  addButton: {
    height: 50,
    borderRadius: 8,
    justifyContent: "center",
    alignItems: "center",
    marginTop: 8,
  },
  addButtonText: {
    color: "white",
    fontSize: 16,
    fontWeight: "600",
  },
  itemsTitle: {
    fontSize: 18,
    fontWeight: "600",
    marginBottom: 12,
    marginTop: 8,
  },
  helperText: {
    fontSize: 14,
    opacity: 0.6,
    textAlign: "center",
    marginTop: 16,
    marginBottom: 16,
  },
});
