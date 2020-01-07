package fr.smartpark.navigator.data

import androidx.room.TypeConverter
import fr.smartpark.navigator.data.models.Point

class Converters {
    @TypeConverter
    fun fromPoint(point: Point) = "${point.lat},${point.long}"

    @TypeConverter
    fun stringToPoint(string: String): Point? {
        val parts = string.split(',', limit = 3)
        if (parts.size != 2) return null
        val t = parts.mapNotNull { it.toFloatOrNull() }
        if (t.size != 2) return null
        return Point(t[0], t[1])
    }
}