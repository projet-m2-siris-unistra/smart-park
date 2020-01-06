package fr.smartpark.navigator

import android.os.Bundle
import android.view.LayoutInflater
import android.view.View
import android.view.ViewGroup
import androidx.navigation.fragment.findNavController
import dagger.android.support.DaggerFragment
import fr.smartpark.navigator.databinding.FragmentTenantListBinding

class TenantListFragment : DaggerFragment() {
    lateinit var binding: FragmentTenantListBinding
    override fun onCreateView(
        inflater: LayoutInflater,
        container: ViewGroup?,
        savedInstanceState: Bundle?
    ): View? {
        binding = FragmentTenantListBinding.inflate(inflater, container, false)
        binding.button.setOnClickListener {
            val direction = TenantListFragmentDirections.actionTenantListToZones(42)
            findNavController().navigate(direction)
        }
        return binding.root
    }
}