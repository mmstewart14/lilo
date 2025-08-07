/**
 * OutfitCard component for displaying outfit recommendations
 */
import { Image } from "expo-image";
import React from "react";
import { StyleSheet, Text, TouchableOpacity, View } from "react-native";
import { useThemeColor } from "../../hooks/useThemeColor";
import { Outfit } from "../../types";
import { ThemedText } from "../ThemedText";

interface OutfitCardProps {
  outfit: Outfit;
  onPress?: () => void;
  onLike?: () => void;
  onDislike?: () => void;
  showActions?: boolean;
}

export function OutfitCard({
  outfit,
  onPress,
  onLike,
  onDislike,
  showActions = true,
}: OutfitCardProps) {
  const backgroundColor = useThemeColor(
    { light: "#fff", dark: "#1c1c1e" },
    "background"
  );
  const textColor = useThemeColor({ light: "#000", dark: "#fff" }, "text");
  const borderColor = useThemeColor(
    { light: "#e0e0e0", dark: "#2c2c2e" },
    "border"
  );

  return (
    <TouchableOpacity
      style={[styles.container, { backgroundColor, borderColor }]}
      onPress={onPress}
      activeOpacity={0.8}
    >
      <View style={styles.imageContainer}>
        {outfit.imageUrl ? (
          <Image
            source={{ uri: outfit.imageUrl }}
            style={styles.image}
            contentFit="cover"
            transition={200}
          />
        ) : (
          <View
            style={[styles.placeholderImage, { backgroundColor: borderColor }]}
          >
            <ThemedText style={styles.placeholderText}>No Image</ThemedText>
          </View>
        )}
        {outfit.isFavorite && (
          <View style={styles.favoriteTag}>
            <ThemedText style={styles.favoriteText}>❤️ Favorite</ThemedText>
          </View>
        )}
      </View>

      <View style={styles.contentContainer}>
        <ThemedText style={styles.title}>{outfit.name}</ThemedText>
        {outfit.description && (
          <ThemedText style={styles.description}>
            {outfit.description}
          </ThemedText>
        )}

        <View style={styles.tagsContainer}>
          {outfit.occasion.map((tag, index) => (
            <View
              key={index}
              style={[styles.tag, { backgroundColor: borderColor }]}
            >
              <Text style={[styles.tagText, { color: textColor }]}>{tag}</Text>
            </View>
          ))}
          {outfit.season.map((tag, index) => (
            <View
              key={`season-${index}`}
              style={[styles.tag, { backgroundColor: borderColor }]}
            >
              <Text style={[styles.tagText, { color: textColor }]}>{tag}</Text>
            </View>
          ))}
        </View>

        {showActions && (
          <View style={styles.actionsContainer}>
            <TouchableOpacity
              style={[styles.actionButton, { backgroundColor: "#4CAF50" }]}
              onPress={onLike}
            >
              <Text style={styles.actionButtonText}>👍 Would Wear</Text>
            </TouchableOpacity>
            <TouchableOpacity
              style={[styles.actionButton, { backgroundColor: "#F44336" }]}
              onPress={onDislike}
            >
              <Text style={styles.actionButtonText}>👎 Wouldnt Wear</Text>
            </TouchableOpacity>
          </View>
        )}
      </View>
    </TouchableOpacity>
  );
}

const styles = StyleSheet.create({
  container: {
    borderRadius: 12,
    overflow: "hidden",
    marginVertical: 8,
    borderWidth: 1,
    shadowColor: "#000",
    shadowOffset: { width: 0, height: 2 },
    shadowOpacity: 0.1,
    shadowRadius: 4,
    elevation: 2,
  },
  imageContainer: {
    position: "relative",
    height: 200,
  },
  image: {
    width: "100%",
    height: "100%",
  },
  placeholderImage: {
    width: "100%",
    height: "100%",
    justifyContent: "center",
    alignItems: "center",
  },
  placeholderText: {
    fontSize: 16,
    fontWeight: "500",
  },
  favoriteTag: {
    position: "absolute",
    top: 10,
    right: 10,
    backgroundColor: "rgba(0, 0, 0, 0.6)",
    paddingHorizontal: 8,
    paddingVertical: 4,
    borderRadius: 4,
  },
  favoriteText: {
    color: "white",
    fontSize: 12,
    fontWeight: "600",
  },
  contentContainer: {
    padding: 16,
  },
  title: {
    fontSize: 18,
    fontWeight: "600",
    marginBottom: 4,
  },
  description: {
    fontSize: 14,
    marginBottom: 12,
    opacity: 0.8,
  },
  tagsContainer: {
    flexDirection: "row",
    flexWrap: "wrap",
    marginBottom: 16,
  },
  tag: {
    paddingHorizontal: 8,
    paddingVertical: 4,
    borderRadius: 4,
    marginRight: 8,
    marginBottom: 8,
  },
  tagText: {
    fontSize: 12,
    fontWeight: "500",
  },
  actionsContainer: {
    flexDirection: "row",
    justifyContent: "space-between",
  },
  actionButton: {
    flex: 1,
    paddingVertical: 10,
    borderRadius: 8,
    justifyContent: "center",
    alignItems: "center",
    marginHorizontal: 4,
  },
  actionButtonText: {
    color: "white",
    fontWeight: "600",
    fontSize: 14,
  },
});
