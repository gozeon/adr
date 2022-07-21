import { reactive, readonly } from "vue";

const modal = reactive<any>({
    role: []
})

export default function useToggleModal() {
    const closeModal = (role = "") => {
        modal.role = modal.role.filter((currentRole: any) => currentRole.type !== role)
    }

    const openModal = (role = "") => {
        modal.role.push({ type: role, isOpen: true })
    }

    const hasRole = (role = "") => {
        if (role === "") {
            return false
        }

        const findRole = modal.role.find((currentRole: any) => currentRole.type === "" ? null : currentRole.type === role)


        return (findRole?.type === role && findRole?.isOpen === true) ? true : false
    }

    const toggleModal = (role = "") => {
        hasRole(role) ? closeModal(role) : openModal(role)
    }

    return {
        state: readonly(modal),
        closeModal,
        openModal,
        hasRole
    }
}