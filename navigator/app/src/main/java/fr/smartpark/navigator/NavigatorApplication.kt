package fr.smartpark.navigator

import android.content.SharedPreferences
import android.util.Log
import androidx.appcompat.app.AppCompatDelegate
import androidx.preference.PreferenceManager
import com.google.android.libraries.maps.MapsInitializer
import dagger.android.AndroidInjector
import dagger.android.support.DaggerApplication
import fr.smartpark.navigator.inject.DaggerAppComponent

class NavigatorApplication : DaggerApplication(), SharedPreferences.OnSharedPreferenceChangeListener {
    override fun applicationInjector(): AndroidInjector<out DaggerApplication> {
        return DaggerAppComponent.factory().create(applicationContext)
    }

    override fun onCreate() {
        super.onCreate()
        MapsInitializer.initialize(this)
        PreferenceManager.getDefaultSharedPreferences(this).let {
            it.registerOnSharedPreferenceChangeListener(this)
            AppCompatDelegate.setDefaultNightMode(
                nightModeFromKey(it.getString("night_mode", null)))
        }
    }

    private fun nightModeFromKey(key: String?) = when(key) {
        "YES" -> AppCompatDelegate.MODE_NIGHT_YES
        "NO" -> AppCompatDelegate.MODE_NIGHT_NO
        else -> AppCompatDelegate.MODE_NIGHT_FOLLOW_SYSTEM
    }

    override fun onSharedPreferenceChanged(sharedPreferences: SharedPreferences?, key: String?) {
        when (key) {
            "night_mode" -> AppCompatDelegate.setDefaultNightMode(
                nightModeFromKey(sharedPreferences?.getString(key, null)))
        }
    }
}