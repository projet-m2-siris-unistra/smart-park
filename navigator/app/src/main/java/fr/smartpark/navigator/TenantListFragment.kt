package fr.smartpark.navigator

import android.content.Intent
import android.net.Uri
import android.os.Bundle
import android.view.*
import androidx.core.content.pm.ShortcutInfoCompat
import androidx.core.content.pm.ShortcutManagerCompat.addDynamicShortcuts
import androidx.core.content.pm.ShortcutManagerCompat.removeAllDynamicShortcuts
import androidx.core.view.updatePadding
import androidx.fragment.app.viewModels
import androidx.lifecycle.ViewModelProvider
import androidx.lifecycle.observe
import androidx.navigation.fragment.findNavController
import dagger.android.support.DaggerFragment
import fr.smartpark.navigator.adapters.TenantAdapter
import fr.smartpark.navigator.api.ApiResult
import fr.smartpark.navigator.data.models.Tenant
import fr.smartpark.navigator.databinding.FragmentTenantListBinding
import fr.smartpark.navigator.viewmodels.TenantListViewModel
import javax.inject.Inject

class TenantListFragment : DaggerFragment() {
    @Inject
    lateinit var viewModelFactory: ViewModelProvider.Factory
    private val viewModel: TenantListViewModel by viewModels { viewModelFactory }

    private lateinit var binding: FragmentTenantListBinding

    override fun onCreateView(
        inflater: LayoutInflater,
        container: ViewGroup?,
        savedInstanceState: Bundle?
    ): View? {
        binding = FragmentTenantListBinding.inflate(inflater, container, false)
        setHasOptionsMenu(true)

        val adapter = TenantAdapter()
        binding.tenantList.adapter = adapter
        viewModel.tenants.observe(viewLifecycleOwner) {
            if (it.status == ApiResult.Status.SUCCESS || it.status == ApiResult.Status.CACHED && !it.data.isNullOrEmpty()) {
                binding.progressBar.visibility = View.GONE
                binding.tenantList.visibility = View.VISIBLE
                adapter.submitList(it.data)
                updateShortcuts(it.data.orEmpty())
            }
        }

        binding.root.setOnApplyWindowInsetsListener { _, insets ->
            binding.tenantList.updatePadding(bottom = insets.systemWindowInsetBottom)
            insets
        }

        return binding.root
    }

    private fun updateShortcuts(tenantList: List<Tenant>) {
        val shortcuts = tenantList.take(4).map {
            ShortcutInfoCompat.Builder(requireContext(), "tenant${it.tenantId}")
                .setShortLabel(it.name)
                .setLongLabel("Show tenant ${it.name}")
                .setIntent(Intent(Intent.ACTION_VIEW, Uri.parse("smartpark://zones?tenant_id=${it.tenantId}")))
                .build()
        }
        removeAllDynamicShortcuts(requireContext())
        addDynamicShortcuts(requireContext(), shortcuts)
    }

    override fun onCreateOptionsMenu(menu: Menu, inflater: MenuInflater) {
        activity?.menuInflater?.inflate(R.menu.menu_main, menu)
        super.onCreateOptionsMenu(menu, inflater)
    }

    override fun onOptionsItemSelected(item: MenuItem): Boolean {
        when (item.itemId) {
            R.id.show_settings ->
                findNavController().navigate(TenantListFragmentDirections.actionTenantListToSettings())
        }
        return super.onOptionsItemSelected(item)
    }

}