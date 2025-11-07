// Last updated: 11/7/2025, 2:49:49 PM
/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     ListNode *next;
 *     ListNode() : val(0), next(nullptr) {}
 *     ListNode(int x) : val(x), next(nullptr) {}
 *     ListNode(int x, ListNode *next) : val(x), next(next) {}
 * };
 */
class Solution {
public:
    ListNode* addTwoNumbers(ListNode* l1, ListNode* l2) {
    ListNode* res = new ListNode(); // Инициализируем указатель на новый узел
        ListNode* current1 = l1;
        ListNode* current2 = l2;
        ListNode* currentRes = res; // Используем отдельный указатель для результата

        int carry = 0; // Переменная для хранения переноса

        while (current1 != nullptr || current2 != nullptr) {
            int x = (current1 != nullptr) ? current1->val : 0;
            int y = (current2 != nullptr) ? current2->val : 0;
            int sum = x + y + carry;
            carry = sum / 10;
            currentRes->next = new ListNode(sum % 10);
            if (current1 != nullptr) current1 = current1->next;
            if (current2 != nullptr) current2 = current2->next;
            currentRes = currentRes->next;
        }
        if (carry > 0) {
            currentRes->next = new ListNode(carry);
        }
        return res->next; 
    }
};