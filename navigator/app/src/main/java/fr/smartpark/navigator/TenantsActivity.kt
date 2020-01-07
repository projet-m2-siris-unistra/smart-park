package fr.smartpark.navigator

import android.os.Bundle
import androidx.appcompat.app.AppCompatActivity
import androidx.databinding.DataBindingUtil
import dagger.android.support.DaggerAppCompatActivity
import fr.smartpark.navigator.databinding.ActivityTenantsBinding

class TenantsActivity : DaggerAppCompatActivity() {
    private lateinit var binding: ActivityTenantsBinding

    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        binding = DataBindingUtil.setContentView(this, R.layout.activity_tenants)
        setSupportActionBar(binding.tenantsToolbar)
    }
}
